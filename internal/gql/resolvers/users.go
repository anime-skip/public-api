package resolvers

import (
	"context"
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, userID *string) (*models.User, error) {
	return &models.User{
		ID:           "ec17af15-e354-440c-a09f-69715fc8b595",
		Email:        "your@email.com",
		CreatedAt:    time.Now(),
		DeletedAt:    nil,
		Username:     "example",
		ProfileURL:   "https://avatars3.githubusercontent.com/u/10101283?s=400&u=2f2037a55606ae978d3088de69584af3586a70cf&v=4",
		AdminOfShows: []*models.ShowAdmin{},
	}, nil
}
