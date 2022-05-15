package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/errors"
)

func ToColorThemeEnum(i int) internal.ColorTheme {
	switch i {
	case internal.THEME_PER_SERVICE:
		return internal.ColorThemePerService
	case internal.THEME_ANIME_SKIP_BLUE:
		return internal.ColorThemeAnimeSkipBlue
	case internal.THEME_VRV_YELLOW:
		return internal.ColorThemeVrvYellow
	case internal.THEME_FUNIMATION_PURPLE:
		return internal.ColorThemeFunimationPurple
	case internal.THEME_CRUNCHYROLL_ORANGE:
		return internal.ColorThemeCrunchyrollOrange
	}
	panic(errors.NewPanicedError("Unknown role integer: %d", i))
}

func ToColorThemeInt(theme internal.ColorTheme) int {
	switch theme {
	case internal.ColorThemePerService:
		return internal.THEME_PER_SERVICE
	case internal.ColorThemeAnimeSkipBlue:
		return internal.THEME_ANIME_SKIP_BLUE
	case internal.ColorThemeVrvYellow:
		return internal.THEME_VRV_YELLOW
	case internal.ColorThemeFunimationPurple:
		return internal.THEME_FUNIMATION_PURPLE
	case internal.ColorThemeCrunchyrollOrange:
		return internal.THEME_CRUNCHYROLL_ORANGE
	}
	panic(errors.NewPanicedError("Unknown theme enum: %s", theme))
}
