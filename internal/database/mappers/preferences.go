package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func DefaultPreferences(userId uuid.UUID) *entities.Preferences {
	return &entities.Preferences{
		UserID: userId,

		EnableAutoSkip:             true,
		EnableAutoPlay:             true,
		MinimizeToolbarWhenEditing: false,
		HideTimelineWhenMinimized:  false,

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

		UserID:                     entity.UserID.String(),
		EnableAutoSkip:             entity.EnableAutoSkip,
		EnableAutoPlay:             entity.EnableAutoPlay,
		MinimizeToolbarWhenEditing: entity.MinimizeToolbarWhenEditing,
		HideTimelineWhenMinimized:  entity.HideTimelineWhenMinimized,

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

	if model.EnableAutoSkip != nil {
		entity.EnableAutoSkip = *model.EnableAutoSkip
	}
	if model.EnableAutoPlay != nil {
		entity.EnableAutoPlay = *model.EnableAutoPlay
	}
	if model.MinimizeToolbarWhenEditing != nil {
		entity.MinimizeToolbarWhenEditing = *model.MinimizeToolbarWhenEditing
	}
	if model.HideTimelineWhenMinimized != nil {
		entity.HideTimelineWhenMinimized = *model.HideTimelineWhenMinimized
	}

	if model.SkipBranding != nil {
		entity.SkipBranding = *model.SkipBranding
	}
	if model.SkipIntros != nil {
		entity.SkipIntros = *model.SkipIntros
	}
	if model.SkipNewIntros != nil {
		entity.SkipNewIntros = *model.SkipNewIntros
	}
	if model.SkipMixedIntros != nil {
		entity.SkipMixedIntros = *model.SkipMixedIntros
	}
	if model.SkipRecaps != nil {
		entity.SkipRecaps = *model.SkipRecaps
	}
	if model.SkipFiller != nil {
		entity.SkipFiller = *model.SkipFiller
	}
	if model.SkipCanon != nil {
		entity.SkipCanon = *model.SkipCanon
	}
	if model.SkipTransitions != nil {
		entity.SkipTransitions = *model.SkipTransitions
	}
	if model.SkipCredits != nil {
		entity.SkipCredits = *model.SkipCredits
	}
	if model.SkipNewCredits != nil {
		entity.SkipNewCredits = *model.SkipNewCredits
	}
	if model.SkipMixedCredits != nil {
		entity.SkipMixedCredits = *model.SkipMixedCredits
	}
	if model.SkipPreview != nil {
		entity.SkipPreview = *model.SkipPreview
	}
	if model.SkipTitleCard != nil {
		entity.SkipTitleCard = *model.SkipTitleCard
	}

	return entity
}
