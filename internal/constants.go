package internal

import (
	"context"

	"github.com/gofrs/uuid"
)

var (
	TIMESTAMP_ID_CANON   = uuid.FromStringOrNil("9edc0037-fa4e-47a7-a29a-d9c43368daa8")
	TIMESTAMP_ID_INTRO   = uuid.FromStringOrNil("14550023-2589-46f0-bfb4-152976506b4c")
	TIMESTAMP_ID_RECAP   = uuid.FromStringOrNil("f38ac196-0d49-40a9-8fcf-f3ef2f40f127")
	TIMESTAMP_ID_CREDITS = uuid.FromStringOrNil("2a730a51-a601-439b-bc1f-7b94a640ffb9")
	TIMESTAMP_ID_PREVIEW = uuid.FromStringOrNil("c7b1eddb-defa-4bc6-a598-f143081cfe4b")
	TIMESTAMP_ID_UNKNOWN = uuid.FromStringOrNil("ae57fcf9-27b0-49a7-9a99-a91aa7518a29")
)

const (
	SHARED_CLIENT_ID = "ZGfO0sMF3eCwLYf8yMSCJjlynwNGRXWE"
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
