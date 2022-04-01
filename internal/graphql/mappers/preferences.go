package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlPreferences(prefs internal.Preferences) graphql.Preferences {
	return graphql.Preferences{
		ID:        &prefs.ID,
		CreatedAt: prefs.CreatedAt,
		UpdatedAt: prefs.UpdatedAt,
		DeletedAt: prefs.DeletedAt,

		UserID:                     &prefs.UserID,
		EnableAutoSkip:             prefs.EnableAutoSkip,
		EnableAutoPlay:             prefs.EnableAutoPlay,
		MinimizeToolbarWhenEditing: prefs.MinimizeToolbarWhenEditing,
		HideTimelineWhenMinimized:  prefs.HideTimelineWhenMinimized,
		ColorTheme:                 ToColorThemeEnum(prefs.ColorTheme),

		SkipBranding:     prefs.SkipBranding,
		SkipIntros:       prefs.SkipIntros,
		SkipNewIntros:    prefs.SkipNewIntros,
		SkipMixedIntros:  prefs.SkipMixedIntros,
		SkipRecaps:       prefs.SkipRecaps,
		SkipFiller:       prefs.SkipFiller,
		SkipCanon:        prefs.SkipCanon,
		SkipTransitions:  prefs.SkipTransitions,
		SkipCredits:      prefs.SkipCredits,
		SkipNewCredits:   prefs.SkipNewCredits,
		SkipMixedCredits: prefs.SkipMixedCredits,
		SkipPreview:      prefs.SkipPreview,
		SkipTitleCard:    prefs.SkipTitleCard,
	}
}

func ToInternalPreferences(prefs graphql.Preferences) internal.Preferences {
	return internal.Preferences{
		ID:        *prefs.ID,
		CreatedAt: prefs.CreatedAt,
		UpdatedAt: prefs.UpdatedAt,
		DeletedAt: prefs.DeletedAt,

		UserID:                     *prefs.UserID,
		EnableAutoSkip:             prefs.EnableAutoSkip,
		EnableAutoPlay:             prefs.EnableAutoPlay,
		MinimizeToolbarWhenEditing: prefs.MinimizeToolbarWhenEditing,
		HideTimelineWhenMinimized:  prefs.HideTimelineWhenMinimized,
		ColorTheme:                 ToColorThemeInt(prefs.ColorTheme),

		SkipBranding:     prefs.SkipBranding,
		SkipIntros:       prefs.SkipIntros,
		SkipNewIntros:    prefs.SkipNewIntros,
		SkipMixedIntros:  prefs.SkipMixedIntros,
		SkipRecaps:       prefs.SkipRecaps,
		SkipFiller:       prefs.SkipFiller,
		SkipCanon:        prefs.SkipCanon,
		SkipTransitions:  prefs.SkipTransitions,
		SkipCredits:      prefs.SkipCredits,
		SkipNewCredits:   prefs.SkipNewCredits,
		SkipMixedCredits: prefs.SkipMixedCredits,
		SkipPreview:      prefs.SkipPreview,
		SkipTitleCard:    prefs.SkipTitleCard,
	}
}
