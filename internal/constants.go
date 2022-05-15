package internal

import (
	"context"

	"github.com/gofrs/uuid"
)

var (
	TIMESTAMP_ID_CANON = uuid.FromStringOrNil("9edc0037-fa4e-47a7-a29a-d9c43368daa8")
	// TIMESTAMP_ID_MUST_WATCH    = "e384759b-3cd2-4824-9569-128363b4452b"
	// TIMESTAMP_ID_BRANDING      = "97e3629a-95e5-4b1a-9411-73a47c0d0e25"
	TIMESTAMP_ID_INTRO = uuid.FromStringOrNil("14550023-2589-46f0-bfb4-152976506b4c")
	// TIMESTAMP_ID_INTRO_INTRO   = "cbb42238-d285-4c88-9e91-feab4bb8ae0a"
	// TIMESTAMP_ID_NEW_INTRO     = "679fb610-ff3c-4cf4-83c0-75bcc7fe8778"
	TIMESTAMP_ID_RECAP = uuid.FromStringOrNil("f38ac196-0d49-40a9-8fcf-f3ef2f40f127")
	// TIMESTAMP_ID_FILLER        = "c48f1dce-1890-4394-8ce6-c3f5b2f95e5e"
	// TIMESTAMP_ID_TRANSITION    = "9f0c6532-ccae-4238-83ec-a2804fe5f7b0"
	TIMESTAMP_ID_CREDITS = uuid.FromStringOrNil("2a730a51-a601-439b-bc1f-7b94a640ffb9")
	// TIMESTAMP_ID_MIXED_CREDITS = "6c4ade53-4fee-447f-89e4-3bb29184e87a"
	// TIMESTAMP_ID_NEW_CREDITS   = "d839cdb1-21b3-455d-9c21-7ffeb37adbec"
	TIMESTAMP_ID_PREVIEW = uuid.FromStringOrNil("c7b1eddb-defa-4bc6-a598-f143081cfe4b")
	// TIMESTAMP_ID_TITLE_CARD    = "67321535-a4ea-4f21-8bed-fb3c8286b510"
	TIMESTAMP_ID_UNKNOWN = uuid.FromStringOrNil("ae57fcf9-27b0-49a7-9a99-a91aa7518a29")
)

var (
	ZeroApiStatus              = ApiStatus{}
	ZeroAPIClient              = APIClient{}
	ZeroAccount                = Account{}
	ZeroEpisode                = Episode{}
	ZeroEpisodeURL             = EpisodeURL{}
	ZeroInputEpisode           = InputEpisode{}
	ZeroInputEpisodeURL        = InputEpisodeURL{}
	ZeroInputExistingTimestamp = InputExistingTimestamp{}
	ZeroInputShow              = InputShow{}
	ZeroInputShowAdmin         = InputShowAdmin{}
	ZeroInputTemplate          = InputTemplate{}
	ZeroInputTemplateTimestamp = InputTemplateTimestamp{}
	ZeroInputTimestamp         = InputTimestamp{}
	ZeroInputTimestampOn       = InputTimestampOn{}
	ZeroInputTimestampType     = InputTimestampType{}
	ZeroLoginData              = LoginData{}
	ZeroPreferences            = Preferences{}
	ZeroShow                   = Show{}
	ZeroShowAdmin              = ShowAdmin{}
	ZeroTemplate               = Template{}
	ZeroTemplateTimestamp      = TemplateTimestamp{}
	ZeroThirdPartyEpisode      = ThirdPartyEpisode{}
	ZeroThirdPartyShow         = ThirdPartyShow{}
	ZeroThirdPartyTimestamp    = ThirdPartyTimestamp{}
	ZeroTimestamp              = Timestamp{}
	ZeroTimestampType          = TimestampType{}
	ZeroUpdatedTimestamps      = UpdatedTimestamps{}
	ZeroUser                   = User{}
	ZeroFullUser               = FullUser{}
)

func NewPreferences(ctx context.Context, userID uuid.UUID) Preferences {
	return Preferences{
		UserID: &userID,

		EnableAutoSkip:             true,
		EnableAutoPlay:             true,
		MinimizeToolbarWhenEditing: false,
		HideTimelineWhenMinimized:  false,
		ColorTheme:                 ColorThemeAnimeSkipBlue,

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
