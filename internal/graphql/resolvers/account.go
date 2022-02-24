package resolvers

import (
	go_context "context"
	"fmt"
	"strings"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/errors"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"anime-skip.com/timestamps-service/internal/log"
	"anime-skip.com/timestamps-service/internal/utils"
	"anime-skip.com/timestamps-service/internal/validation"
)

// Helpers

func (r *Resolver) getLoginData(ctx go_context.Context, user internal.User) (*graphql.LoginData, error) {
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

	account := mappers.ToGraphqlAccount(user)
	return &graphql.LoginData{
		AuthToken:    accessToken,
		RefreshToken: refreshToken,
		Account:      &account,
	}, nil
}

// Mutations

func (r *mutationResolver) CreateAccount(ctx go_context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*graphql.LoginData, error) {
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
	_, err = r.UserService.GetByUsername(ctx, username)
	if err == nil {
		return nil, fmt.Errorf("username='%s' is already taken, use a different one", username)
	}
	if !errors.IsRecordNotFound(err) {
		return nil, fmt.Errorf("Error checking for user with same username: %s", err.Error())
	}

	log.V("Checking for existing email")
	_, err = r.UserService.GetByEmail(ctx, email)
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

	tx := r.DB.MustBeginTx(ctx, nil)
	defer tx.Rollback()

	log.V("Creating user")
	createdUser, err := r.UserService.CreateInTx(ctx, tx, internal.User{
		ID:            utils.RandomID(),
		Username:      username,
		Email:         email,
		PasswordHash:  encryptedPasswordHash,
		EmailVerified: false,
		Role:          internal.ROLE_USER,
	})
	if err != nil {
		return nil, err
	}

	log.V("Creating Preferences")
	defaultPreferences := r.PreferencesService.NewDefault(ctx, createdUser.ID)
	_, err = r.PreferencesService.CreateInTx(ctx, tx, defaultPreferences)
	if err != nil {
		return nil, err
	}

	log.V("Sending welcome email")
	err = r.EmailService.SendWelcome(createdUser)
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

	log.V("Commiting transaction")
	tx.Commit()
	account := mappers.ToGraphqlAccount(createdUser)

	log.V("Creating email token")
	verifyEmailToken, err := r.AuthService.CreateVerifyEmailToken(createdUser)
	if err != nil {
		log.E("Failed to create verify email token: %v", err)
	} else {
		err = r.EmailService.SendVerification(createdUser, verifyEmailToken)
		if err != nil {
			log.E("Failed to send email address verification email (but still created user): %v", err)
		}
	}

	log.V("Returning LoginData")
	return &graphql.LoginData{
		AuthToken:    accessToken,
		RefreshToken: refreshToken,
		Account:      &account,
	}, nil
}

func (r *mutationResolver) ChangePassword(ctx go_context.Context, oldPassword string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("mutationResolver.ChangePassword not implemented")
}

func (r *mutationResolver) ResendVerificationEmail(ctx go_context.Context, recaptchaResponse string) (*bool, error) {
	panic("mutationResolver.ResendVerificationEmail not implemented")
}

func (r *mutationResolver) VerifyEmailAddress(ctx go_context.Context, validationToken string) (*graphql.Account, error) {
	panic("mutationResolver.VerifyEmailAddress not implemented")
}

func (r *mutationResolver) RequestPasswordReset(ctx go_context.Context, recaptchaResponse string, email string) (bool, error) {
	panic("mutationResolver.RequestPasswordReset not implemented")
}

func (r *mutationResolver) ResetPassword(ctx go_context.Context, passwordResetToken string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("mutationResolver.ResetPassword not implemented")
}

func (r *mutationResolver) DeleteAccountRequest(ctx go_context.Context, passwordHash string) (*graphql.Account, error) {
	return nil, fmt.Errorf("mutationResolver.DeleteAccountRequest not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx go_context.Context, deleteToken string) (*graphql.Account, error) {
	return nil, fmt.Errorf("mutationResolver.DeleteAccount not implemented")
}

// Queries

func (r *queryResolver) Login(ctx go_context.Context, usernameOrEmail string, passwordHash string) (*graphql.LoginData, error) {
	usernameOrEmail = strings.TrimSpace(usernameOrEmail)
	passwordHash = strings.TrimSpace(passwordHash)

	user, err := r.UserService.GetByUsernameOrEmail(ctx, usernameOrEmail)
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

func (r *queryResolver) LoginRefresh(ctx go_context.Context, refreshToken string) (*graphql.LoginData, error) {
	claims, err := r.AuthService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	user, err := r.UserService.GetByID(ctx, claims.UserID)
	if err != nil {
		log.V("Failed to get user with id='%s': %v", claims.UserID, err)
		return nil, fmt.Errorf("Bad login credentials")
	}

	return r.getLoginData(ctx, user)
}

// Fields

func (r *queryResolver) Account(ctx go_context.Context) (*graphql.Account, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	internalUser, err := r.UserService.GetByID(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	account := mappers.ToGraphqlAccount(internalUser)
	return &account, nil
}

func (r *accountResolver) Preferences(ctx go_context.Context, obj *graphql.Account) (*graphql.Preferences, error) {
	return r.getPreferences(ctx, *obj.ID)
}

func (r *accountResolver) AdminOfShows(ctx go_context.Context, obj *graphql.Account) ([]*graphql.ShowAdmin, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	internalShowAdmins, err := r.ShowAdminService.GetByUserID(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	showAdmins := mappers.ToGraphqlShowAdminPointers(internalShowAdmins)
	return showAdmins, nil
}
