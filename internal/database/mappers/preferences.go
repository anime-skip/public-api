package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// PreferencesEntityToModel -
func PreferencesEntityToModel(entity *entities.Preferences) (*models.Preferences, error) {
	model := &models.Preferences{
		ID:        entity.ID.String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,

		UserID:           entity.UserID.String(),
		EnableAutoSkip:   entity.EnableAutoSkip,
		EnableAutoPlay:   entity.EnableAutoPlay,
		SkipBranding:     entity.SkipBranding,
		SkipIntros:       entity.SkipIntros,
		SkipNewIntros:    entity.SkipNewIntros,
		SkipMixedIntros:  entity.SkipMixedIntros,
		SkipRecaps:       entity.SkipRecaps,
		SkipFiller:       entity.SkipFiller,
		SkipCanon:        entity.SkipCanon,
		SkipTransitions:  entity.SkipTransitions,
		SkipCredits:      entity.SkipCredits,
		SkipNewCredits:   entity.SkipNewCredits,
		SkipMixedCredits: entity.SkipMixedCredits,
		SkipPreview:      entity.SkipPreview,
		SkipTitleCard:    entity.SkipTitleCard,
	}
	return model, nil
}