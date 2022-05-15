package resolvers

import (
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"anime-skip.com/public-api/internal/validation"
)

// Helpers

func (r *Resolver) getLoginData(ctx context.Context, user internal.FullUser) (*internal.LoginData, error) {
	accessToken, err := r.AuthService.CreateAccessToken(user)
	if err != nil {
		log.E("Failed to generate an auth token: %v", err)
		return nil, fmt.Errorf("Failed to login")
	}

	refreshToken, err := r.AuthService.CreateRefreshToken(user)
	if err != nil {
		log.E("Failed to generate a refresh token: %v", err)
		return nil, fmt.Errorf("Failed to login")
	}

	account := mappers.ToAccount(user)
	return &internal.LoginData{
		AuthToken:    accessToken,
		RefreshToken: refreshToken,
		Account:      &account,
	}, nil
}

// Mutations

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*internal.LoginData, error) {
	log.V("Additional input validation")
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	passwordHash = strings.TrimSpace(passwordHash)

	if err := validation.AccountUsername(username); err != nil {
		return nil, err
	}
	if err := validation.AccountEmail(email); err != nil {
		return nil, err
	}

	log.V("Verify recaptcha")
	err := r.RecaptchaService.Verify(ctx, recaptchaResponse)
	if err != nil {
		return nil, err
	}

	log.V("Checking for existing username")
	_, err = r.UserService.Get(ctx, internal.UsersFilter{
		Username: &username,
	})
	if err == nil {
		return nil, fmt.Errorf("username='%s' is already taken, use a different one", username)
	}
	if !errors.IsRecordNotFound(err) {
		return nil, fmt.Errorf("Error checking for user with same username: %s", err.Error())
	}

	log.V("Checking for existing email")
	_, err = r.UserService.Get(ctx, internal.UsersFilter{
		Email: &email,
	})
	if err == nil {
		return nil, fmt.Errorf("email='%s' is already taken, use a different one", email)
	}
	if !errors.IsRecordNotFound(err) {
		return nil, fmt.Errorf("Error checking for user with same email: %s", err.Error())
	}

	log.V("Generating passwordHash")
	encryptedPasswordHash, err := r.AuthService.CreateEncryptedPassword(passwordHash)
	if err != nil {
		return nil, err
	}

	log.V("Creating user and preferences")
	createdUser, err := r.UserService.CreateAccount(ctx, internal.FullUser{
		Username:      username,
		Email:         email,
		PasswordHash:  encryptedPasswordHash,
		EmailVerified: false,
		Role:          internal.ROLE_USER,
	})
	if err != nil {
		return nil, err
	}

	log.V("Sending welcome email")
	err = r.EmailService.SendWelcome(ctx, createdUser)
	if err != nil {
		log.E("Failed to send welcome email: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	log.V("Creating access token")
	accessToken, err := r.AuthService.CreateAccessToken(createdUser)
	if err != nil {
		log.E("Failed to create access token: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	log.V("Creating refresh token")
	refreshToken, err := r.AuthService.CreateRefreshToken(createdUser)
	if err != nil {
		log.E("Failed to create refresh token: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	account := mappers.ToAccount(createdUser)

	log.V("Creating email token")
	verifyEmailToken, err := r.AuthService.CreateVerifyEmailToken(createdUser)
	if err != nil {
		log.E("Failed to create verify email token: %v", err)
	} else {
		err = r.EmailService.SendVerification(ctx, createdUser, verifyEmailToken)
		if err != nil {
			log.E("Failed to send email address verification email (but still created user): %v", err)
		}
	}

	log.V("Returning LoginData")
	return &internal.LoginData{
		AuthToken:    accessToken,
		RefreshToken: refreshToken,
		Account:      &account,
	}, nil
}

func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string, confirmNewPassword string) (*internal.LoginData, error) {
	if newPassword != confirmNewPassword {
		return nil, errors.New("New passwords do not match")
	}
	if newPassword == "" {
		return nil, errors.New("New password is not valid, it cannot be empty")
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}

	oldPasswordHash := utils.MD5(oldPassword)
	if err = r.AuthService.ValidatePassword(oldPasswordHash, user.PasswordHash); err != nil {
		return nil, fmt.Errorf("Old password is not correct")
	}

	newPasswordHash := utils.MD5(newPassword)
	newEncryptedPasswordHash, err := r.AuthService.CreateEncryptedPassword(newPasswordHash)
	if err != nil {
		return nil, err
	}

	userWithNewPassword := user
	userWithNewPassword.PasswordHash = newEncryptedPasswordHash
	newUser, err := r.UserService.Update(ctx, userWithNewPassword)
	if err != nil {
		return nil, err
	}

	return r.getLoginData(ctx, newUser)
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context, recaptchaResponse string) (*bool, error) {
	err := r.RecaptchaService.Verify(ctx, recaptchaResponse)
	if err != nil {
		return nil, err
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}
	token, err := r.AuthService.CreateVerifyEmailToken(user)
	if err != nil {
		return nil, err
	}

	err = r.EmailService.SendVerification(ctx, user, token)
	isSent := err == nil
	return &isSent, err
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*internal.Account, error) {
	claims, err := r.AuthService.ValidateVerifyEmailToken(validationToken)
	if err != nil {
		return nil, err
	}

	// Update the user to have their email verified
	existingUser, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &claims.UserID,
	})
	if err != nil {
		return nil, err
	}
	existingUser.EmailVerified = true
	updatedUser, err := r.UserService.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	account := mappers.ToAccount(updatedUser)
	return &account, nil
}

func (r *mutationResolver) RequestPasswordReset(ctx context.Context, recaptchaResponse string, email string) (bool, error) {
	email = strings.TrimSpace(email)
	err := validation.AccountEmail(email)
	if err != nil {
		return false, err
	}
	err = r.RecaptchaService.Verify(ctx, recaptchaResponse)
	if err != nil {
		return false, err
	}

	user, err := r.UserService.Get(ctx, internal.UsersFilter{
		Email: &email,
	})
	if errors.IsRecordNotFound(err) {
		// Don't provide hints to if a user has an account or not
		return true, nil
	} else if err != nil {
		return false, err
	}

	token, err := r.AuthService.CreateResetPasswordToken(user)
	if err != nil {
		return false, err
	}
	err = r.EmailService.SendResetPassword(ctx, user, token)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, passwordResetToken string, newPassword string, confirmNewPassword string) (*internal.LoginData, error) {
	if newPassword != confirmNewPassword {
		return nil, errors.New("New passwords don't match")
	}

	claims, err := r.AuthService.ValidateResetPasswordToken(passwordResetToken)
	if err != nil {
		return nil, err
	}

	newPasswordHash := utils.MD5(newPassword)
	newEncryptedPasswordHash, err := r.AuthService.CreateEncryptedPassword(newPasswordHash)
	if err != nil {
		return nil, err
	}

	existingUser, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &claims.UserID,
	})
	if err != nil {
		return nil, err
	}
	existingUser.PasswordHash = newEncryptedPasswordHash

	newUser, err := r.UserService.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return r.getLoginData(ctx, newUser)
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, passwordHash string) (*internal.Account, error) {
	return nil, fmt.Errorf("TODO - currently the api doesn't support deleting accounts")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, deleteToken string) (*internal.Account, error) {
	return nil, fmt.Errorf("TODO - currently the api doesn't support deleting accounts")
}

// Queries

func (r *queryResolver) Login(ctx context.Context, usernameOrEmail string, passwordHash string) (*internal.LoginData, error) {
	usernameOrEmail = strings.TrimSpace(usernameOrEmail)
	passwordHash = strings.TrimSpace(passwordHash)

	user, err := r.UserService.Get(ctx, internal.UsersFilter{
		UsernameOrEmail: &usernameOrEmail,
	})
	if err != nil {
		log.V("Failed to get user for username or email = '%s': %v", usernameOrEmail, err)
		// auth.LoginRetryTimer.Failure(usernameOrEmail)
		return nil, fmt.Errorf("Bad login credentials")
	}

	if err = r.AuthService.ValidatePassword(passwordHash, user.PasswordHash); err != nil {
		log.V("Failed validate password: %v", err)
		// auth.LoginRetryTimer.Failure(usernameOrEmail)
		return nil, fmt.Errorf("Bad login credentials")
	}

	// defer auth.LoginRetryTimer.Success(usernameOrEmail)
	return r.getLoginData(ctx, user)
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*internal.LoginData, error) {
	claims, err := r.AuthService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	user, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &claims.UserID,
	})
	if err != nil {
		log.V("Failed to get user with id='%s': %v", claims.UserID, err)
		return nil, fmt.Errorf("Bad login credentials")
	}

	return r.getLoginData(ctx, user)
}

// Fields

func (r *queryResolver) Account(ctx context.Context) (*internal.Account, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	internalUser, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}
	account := mappers.ToAccount(internalUser)
	return &account, nil
}

func (r *accountResolver) Preferences(ctx context.Context, obj *internal.Account) (*internal.Preferences, error) {
	return r.getPreferences(ctx, *obj.ID)
}

func (r *accountResolver) AdminOfShows(ctx context.Context, obj *internal.Account) ([]*internal.ShowAdmin, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	showAdmins, err := r.ShowAdminService.List(ctx, internal.ShowAdminsFilter{
		UserID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(showAdmins), nil
}
