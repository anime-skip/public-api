package resolvers

import (
	"context"
	"fmt"

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

// Mutation Resolvers

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string) (*models.Account, error) {
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

	user, err := repos.CreateUser(tx, username, email, encryptedPasswordHash)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = emailService.SendWelcome(user)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Failed to create user")
	}

	utils.CommitTransaction(tx, false)

	validateEmailToken, err := utils.GenerateValidateEmailToken(user)
	if err != nil {
		log.E("Failed to send token validation email: %v", err)
	} else {
		emailService.SendEmailAddressValidation(user, validateEmailToken)
	}

	return mappers.UserEntityToAccountModel(user), nil
}

func (r *mutationResolver) SendEmailAddressValidationEmail(ctx context.Context, userID string) (*bool, error) {
	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateValidateEmailToken(user)
	if err != nil {
		return nil, err
	}

	err = emailService.SendEmailAddressValidation(user, token)
	isSent := err == nil
	return &isSent, nil
}

func (r *mutationResolver) ValidateEmailAddress(ctx context.Context, validationToken string) (*models.Account, error) {
	return nil, fmt.Errorf("not implemented")
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
