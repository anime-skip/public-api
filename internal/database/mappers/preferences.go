package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func DefaultPreferences(userId uuid.UUID) *entities.Preferences {
	return &entities.Preferences{
		UserID:           userId,
		EnableAutoSkip:   true,
		EnableAutoPlay:   true,
		SkipBranding:     true,
		SkipIntros:       true,
		SkipNewIntros:    false,
		SkipMixedIntros:  false,
		SkipRecaps:       true,
		SkipFiller:       true,
		SkipCanon:        false,
		SkipTransitions:  true,
		SkipCredits:      true,
		SkipNewCredits:   false,
		SkipMixedCredits: false,
		SkipPreview:      true,
		SkipTitleCard:    true,
	}
}

// PreferencesEntityToModel -
func PreferencesEntityToModel(entity *entities.Preferences) *models.Preferences {
	if entity == nil {
		return nil
	}
	return &models.Preferences{
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
}

func PreferencesInputModelToEntity(model models.InputPreferences, entity *entities.Preferences) *entities.Preferences {
	if entity == nil {
		return nil
	}
	entity.EnableAutoSkip = model.EnableAutoSkip
	entity.EnableAutoPlay = model.EnableAutoPlay
	entity.SkipBranding = model.SkipBranding
	entity.SkipIntros = model.SkipIntros
	entity.SkipNewIntros = model.SkipNewIntros
	entity.SkipMixedIntros = model.SkipMixedIntros
	entity.SkipRecaps = model.SkipRecaps
	entity.SkipFiller = model.SkipFiller
	entity.SkipCanon = model.SkipCanon
	entity.SkipTransitions = model.SkipTransitions
	entity.SkipCredits = model.SkipCredits
	entity.SkipNewCredits = model.SkipNewCredits
	entity.SkipMixedCredits = model.SkipMixedCredits
	entity.SkipPreview = model.SkipPreview
	entity.SkipTitleCard = model.SkipTitleCard
	return entity
}
