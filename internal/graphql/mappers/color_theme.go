package mappers

import (
	"fmt"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToColorThemeEnum(i int) graphql.ColorTheme {
	switch i {
	case internal.THEME_PER_SERVICE:
		return graphql.ColorThemePerService
	case internal.THEME_ANIME_SKIP_BLUE:
		return graphql.ColorThemeAnimeSkipBlue
	case internal.THEME_VRV_YELLOW:
		return graphql.ColorThemeVrvYellow
	case internal.THEME_FUNIMATION_PURPLE:
		return graphql.ColorThemeFunimationPurple
	case internal.THEME_CRUNCHYROLL_ORANGE:
		return graphql.ColorThemeCrunchyrollOrange
	}
	panic(fmt.Errorf("Unknown role integer: %d", i))
}

func ToColorThemeInt(theme graphql.ColorTheme) int {
	switch theme {
	case graphql.ColorThemePerService:
		return internal.THEME_PER_SERVICE
	case graphql.ColorThemeAnimeSkipBlue:
		return internal.THEME_ANIME_SKIP_BLUE
	case graphql.ColorThemeVrvYellow:
		return internal.THEME_VRV_YELLOW
	case graphql.ColorThemeFunimationPurple:
		return internal.THEME_FUNIMATION_PURPLE
	case graphql.ColorThemeCrunchyrollOrange:
		return internal.THEME_CRUNCHYROLL_ORANGE
	}
	panic(fmt.Errorf("Unknown theme enum: %s", theme))
}
