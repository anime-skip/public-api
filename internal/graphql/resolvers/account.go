package resolvers

import (
	"context"
	"errors"
	"fmt"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	emailService "anime-skip.com/backend/internal/services/email"
	"anime-skip.com/backend/internal/services/recaptcha"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/auth"
	"anime-skip.com/backend/internal/utils/context_utils"
	"anime-skip.com/backend/internal/utils/log"
)

// Query Resolvers

type AccountResolver struct{ *Resolver }

func (r *queryResolver) Account(ctx context.Context) (*models.Account, error) {
	userID, err := context_utils.UserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToAccountModel(user), nil
}

func (r *queryResolver) Login(ctx context.Context, usernameEmail string, passwordHash string) (*models.LoginData, error) {
	user, err := repos.FindUserByUsernameOrEmail(r.DB(ctx), usernameEmail)
	if err != nil {
		log.V("Failed to get user for username or email = '%s': %v", usernameEmail, err)
		auth.LoginRetryTimer.Failure(usernameEmail)
		return nil, fmt.Errorf("Bad login credentials")
	}

	if err = auth.ValidatePassword(passwordHash, user.PasswordHash); err != nil {
		log.V("Failed validate password: %v", err)
		auth.LoginRetryTimer.Failure(usernameEmail)
		return nil, fmt.Errorf("Bad login credentials")
	}

	authToken, err := auth.GenerateAuthToken(user)
	if err != nil {
		log.V("Failed to generate auth token for %v: %v", usernameEmail, err)
		return nil, fmt.Errorf("Failed to login")
	}

	refreshToken, err := auth.GenerateRefreshToken(user)
	if err != nil {
		log.V("Failed to generate auth token for %v: %v", usernameEmail, err)
		return nil, fmt.Errorf("Failed to login")
	}

	defer auth.LoginRetryTimer.Success(usernameEmail)
	return &models.LoginData{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		Account:      mappers.UserEntityToAccountModel(user),
	}, nil
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*models.LoginData, error) {
	claims, err := auth.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	userID := claims["userId"].(string)
	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		log.V("Failed to get user with id='%s': %v", userID, err)
		return nil, fmt.Errorf("Bad login credentials")
	}

	authToken, err := auth.GenerateAuthToken(user)
	if err != nil {
		log.V("Failed to generate auth token: %v", err)
		return nil, fmt.Errorf("Failed to login")
	}

	newRefreshToken, err := auth.GenerateRefreshToken(user)
	if err != nil {
		log.V("Failed to generate auth token: %v", err)
		return nil, fmt.Errorf("Failed to login")
	}

	return &models.LoginData{
		AuthToken:    authToken,
		RefreshToken: newRefreshToken,
		Account:      mappers.UserEntityToAccountModel(user),
	}, nil
}

// Mutation Resolvers

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*models.LoginData, error) {
	tx := utils.StartTransaction(r.DB(ctx), false)

	existingUser, _ := repos.FindUserByUsername(tx, username)
	if existingUser != nil {
		tx.Rollback()
		return nil, fmt.Errorf("username='%s' is already taken, use a different one", username)
	}

	existingUser, _ = repos.FindUserByEmail(tx, email)
	if existingUser != nil {
		tx.Rollback()
		return nil, fmt.Errorf("email='%s' is already taken, use a different one", email)
	}

	encryptedPasswordHash, err := auth.GenerateEncryptedPassword(passwordHash)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	ipAddress, err := context_utils.IPAddress(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Could not get ip address from request")
	}
	err = recaptcha.Verify(recaptchaResponse, ipAddress)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user, err := repos.CreateUser(tx, username, email, encryptedPasswordHash)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = emailService.SendWelcome(user)
	if err != nil {
		tx.Rollback()
		log.E("Failed to send welcome email: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	account := mappers.UserEntityToAccountModel(user)

	authToken, err := auth.GenerateAuthToken(user)
	if err != nil {
		tx.Rollback()
		log.E("Failed to create auth token: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	refreshToken, err := auth.GenerateRefreshToken(user)
	if err != nil {
		tx.Rollback()
		log.E("Failed to create auth token: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	utils.CommitTransaction(tx, false)

	verifyEmailToken, err := auth.GenerateVerifyEmailToken(user)
	if err != nil {
		log.E("Failed to send token validation email: %v", err)
	} else {
		err = emailService.SendVerification(user, verifyEmailToken)
		if err != nil {
			log.E("Failed to send email address verification email (but still created user): %v", err)
		}
	}

	return &models.LoginData{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		Account:      account,
	}, nil
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context) (*bool, error) {
	userID, err := context_utils.UserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateVerifyEmailToken(user)
	if err != nil {
		return nil, err
	}

	err = emailService.SendVerification(user, token)
	isSent := err == nil
	return &isSent, err
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*models.Account, error) {
	payload, err := auth.ValidateEmailVerificationToken(validationToken)
	if err != nil {
		return nil, err
	}

	// Update the user to have their email verified
	userID := payload["userId"].(string)
	existingUser, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}
	updatedUser, err := repos.VerifyUserEmail(r.DB(ctx), existingUser)
	if err != nil {
		return nil, err
	}

	return mappers.UserEntityToAccountModel(updatedUser), nil
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, passwordHash string) (*models.Account, error) {
	userID, err := context_utils.UserID(ctx)
	if err != nil {
		return nil, err
	}
	log.I(userID)

	return nil, fmt.Errorf("not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, deleteToken string) (*models.Account, error) {
	return nil, fmt.Errorf("not implemented")
}

// Field Resolvers

func (r *AccountResolver) AdminOfShows(ctx context.Context, obj *models.Account) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(r.DB(ctx), obj.ID)
}

func (r *AccountResolver) Preferences(ctx context.Context, obj *models.Account) (*models.Preferences, error) {
	preferences, err := repos.FindPreferencesByUserID(r.DB(ctx), obj.ID)
	return mappers.PreferencesEntityToModel(preferences), err
}
