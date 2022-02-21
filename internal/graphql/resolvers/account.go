package resolvers

import (
	go_context "context"

	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
)

// Mutations

func (r *mutationResolver) CreateAccount(ctx go_context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*graphql.LoginData, error) {
	panic("mutationResolver.CreateAccount not implemented")
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
	panic("mutationResolver.DeleteAccountRequest not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx go_context.Context, deleteToken string) (*graphql.Account, error) {
	panic("mutationResolver.DeleteAccount not implemented")
}

// Queries

func (r *queryResolver) Login(ctx go_context.Context, usernameEmail string, passwordHash string) (*graphql.LoginData, error) {
	panic("queryResolver.Login not implemented")
}

func (r *queryResolver) LoginRefresh(ctx go_context.Context, refreshToken string) (*graphql.LoginData, error) {
	panic("queryResolver.LoginRefresh not implemented")
}

// Fields

func (r *queryResolver) Account(ctx go_context.Context) (*graphql.Account, error) {
	auth, err := context.GetAuthenticationDetails(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.UserService.GetByID(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	account := mappers.ToGraphqlAccount(user)
	return &account, nil
}

func (r *accountResolver) Preferences(ctx go_context.Context, obj *graphql.Account) (*graphql.Preferences, error) {
	return r.getPreferences(ctx, *obj.ID)
}

func (r *accountResolver) AdminOfShows(ctx go_context.Context, obj *graphql.Account) ([]*graphql.ShowAdmin, error) {
	panic("accountResolver.AdminOfShows not implemented")
}
