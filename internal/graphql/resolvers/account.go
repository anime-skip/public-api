package resolvers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	emailService "github.com/aklinker1/anime-skip-backend/internal/server/email"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

// Helpers

// Query Resolvers

type AccountResolver struct{ *Resolver }

func (r *queryResolver) Account(ctx context.Context) (*models.Account, error) {
	userID, err := utils.UserIDFromContext(ctx)
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
		time.Sleep(3 * time.Second)
		return nil, fmt.Errorf("Bad login credentials")
	}

	if err = utils.ValidatePassword(passwordHash, user.PasswordHash); err != nil {
		log.V("Failed validate password: %v", err)
		return nil, fmt.Errorf("Bad login credentials")
	}

	authToken, err := utils.GenerateAuthToken(user)
	if err != nil {
		log.V("Failed to generate auth token: %v", usernameEmail, err)
		return nil, fmt.Errorf("Failed to login")
	}

	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		log.V("Failed to generate auth token: %v", usernameEmail, err)
		return nil, fmt.Errorf("Failed to login")
	}

	return &models.LoginData{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		Account:      mappers.UserEntityToAccountModel(user),
	}, nil
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*models.LoginData, error) {
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	userID := claims["userId"].(string)
	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		log.V("Failed to get user with id='%s': %v", userID, err)
		return nil, fmt.Errorf("Bad login credentials")
	}

	authToken, err := utils.GenerateAuthToken(user)
	if err != nil {
		log.V("Failed to generate auth token: %v", err)
		return nil, fmt.Errorf("Failed to login")
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user)
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

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*models.Account, error) {
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

	encryptedPasswordHash, err := utils.GenerateEncryptedPassword(passwordHash)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	ipAddress, err := utils.GetIP(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Could not get ip address from request")
	}
	err = utils.VerifyRecaptcha(recaptchaResponse, ipAddress)
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
		log.E("Failed to send email: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}

	utils.CommitTransaction(tx, false)

	verifyEmailToken, err := utils.GenerateVerifyEmailToken(user)
	if err != nil {
		log.E("Failed to send token validation email: %v", err)
	} else {
		err = emailService.SendVerification(user, verifyEmailToken)
		if err != nil {
			log.E("Failed to send email address verification email: %v", err)
		}
	}

	return mappers.UserEntityToAccountModel(user), nil
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context, userID string) (*bool, error) {
	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateVerifyEmailToken(user)
	if err != nil {
		return nil, err
	}

	err = emailService.SendVerification(user, token)
	isSent := err == nil
	return &isSent, err
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*models.Account, error) {
	payload, err := utils.ValidateEmailVerificationToken(validationToken)
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

	return mappers.UserEntityToAccountModel(updatedUser), nil
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, accoutnID string, passwordHash string) (*models.Account, error) {
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
