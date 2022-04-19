// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

// The base model has all the fields you would expect a fully fleshed out item in the database would
// have. It is used to track who create, updated, and deleted items
type BaseModel interface {
	IsBaseModel()
}

// Account info that should only be accessible by the authorised user
type Account struct {
	ID        *uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	// Unique string slug that is the easy to remember identifier
	Username string `json:"username"`
	Email    string `json:"email"`
	// Url to an image that is the user's profile picture
	ProfileURL string `json:"profileUrl"`
	// The linking object that associates a user to the shows they are admins of.
	//
	// > This data is also accessible on the `User` model. It has been added here for convienience
	AdminOfShows []*ShowAdmin `json:"adminOfShows"`
	// If the user's email is verified. Emails must be verified before the user can call a mutation
	EmailVerified bool `json:"emailVerified"`
	// The user's administrative role. Most users are `Role.USER`
	Role Role `json:"role"`
	// The user's preferences
	Preferences *Preferences `json:"preferences"`
}

// Basic information about an episode, including season, numbers, a list of timestamps, and urls that
// it can be watched at
type Episode struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The season number that this episode belongs to
	//
	// ### Examples:
	//
	// - "1"
	// - "1 Directors Cut"
	// - "2"
	// - "Movies"
	Season *string `json:"season"`
	// The episode number in the current season
	//
	// ### Examples:
	//
	// - "1"
	// - "2"
	// - "5.5"
	// - "OVA 1"
	Number *string `json:"number"`
	// The absolute episode number out of all the episodes of the show. Generally only regular episodes
	// should have this field
	AbsoluteNumber *string `json:"absoluteNumber"`
	// The duration of the episode's first url, which can be used to calculate a suggested offset for new
	// episode urls. Episodes at different URLs have different branding intros, and that difference can
	// be computed using: `EpisodeUrl.duration - Episode.baseDuration`
	// Generally, this works because each service has it's own branding at the beginning of the show, not
	// at the end of it
	BaseDuration *float64 `json:"baseDuration"`
	// The episode's name
	Name *string `json:"name"`
	// The show that the episode belongs to
	Show *Show `json:"show"`
	// The id of the show that the episode blongs to
	ShowID *uuid.UUID `json:"showId"`
	// The list of current timestamps.
	//
	// Timestamps are apart apart of the `Episode` instead of the `EpisodeUrl` so that they can be shared
	// between urls and not need duplicate data
	Timestamps []*Timestamp `json:"timestamps"`
	// The list of urls and services that the episode can be accessed from
	Urls []*EpisodeURL `json:"urls"`
	// If the episode is the source episode for a `Template`, this will resolve to that template
	Template *Template `json:"template"`
}

func (Episode) IsBaseModel() {}

// Stores information about what where an episode can be watched from
type EpisodeURL struct {
	// The url that would take a user to watch the `episode`.
	//
	// This url should be stripped of all query params.
	URL             string     `json:"url"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	// The length of the episode at this url. For more information on why this field exists, check out
	// the `Episode.baseDuration`. If an `Episode` does not have a duration, that `Episode` and this
	// `EpisodeUrl` should be given the same value, and the `EpisodeUrl.timestampsOffset` should be set to 0
	Duration *float64 `json:"duration"`
	// How much a episode's timestamps should be offset for this `EpisodeUrl`, since different services
	// have different branding animations, leading to offsets between services. This field can be edited
	// to whatever, but it should be suggested to be `EpisodeUrl.duration - Episode.baseDuration`.
	// It can be positive or negative.
	TimestampsOffset *float64 `json:"timestampsOffset"`
	// The `Episode.id` that this url belongs to
	EpisodeID *uuid.UUID `json:"episodeId"`
	// The `Episode` that this url belongs to
	Episode *Episode `json:"episode"`
	// What service this url points to. This is computed when the `EpisodeUrl` is created
	Source EpisodeSource `json:"source"`
}

// Data required to create a new `Episode`. See `Episode` for a description of each field
type InputEpisode struct {
	// See `Episode.season`
	Season *string `json:"season"`
	// See `Episode.number`
	Number *string `json:"number"`
	// See `Episode.absoluteNumber`
	AbsoluteNumber *string `json:"absoluteNumber"`
	// See `Episode.name`
	Name *string `json:"name"`
	// See `Episode.baseDuration`
	BaseDuration *float64 `json:"baseDuration"`
}

// Data required to create a new `EpisodeUrl`. See `EpisodeUrl` for a description of each field
type InputEpisodeURL struct {
	URL              string   `json:"url"`
	Duration         *float64 `json:"duration"`
	TimestampsOffset *float64 `json:"timestampsOffset"`
}

type InputExistingTimestamp struct {
	// The id of the timestamp you want to modify
	ID *uuid.UUID `json:"id"`
	// The new values for the timestamp
	Timestamp *InputTimestamp `json:"timestamp"`
}

// Data required to create a new `Show`. See `Show` for a description of each field
type InputShow struct {
	Name         string  `json:"name"`
	OriginalName *string `json:"originalName"`
	Website      *string `json:"website"`
	Image        *string `json:"image"`
}

// Data required to create a new `ShowAdmin`. See `ShowAdmin` for a description of each field
type InputShowAdmin struct {
	ShowID *uuid.UUID `json:"showId"`
	UserID *uuid.UUID `json:"userId"`
}

// Data required to create a new template. See `Template` for a description of each field
type InputTemplate struct {
	ShowID          *uuid.UUID   `json:"showId"`
	Type            TemplateType `json:"type"`
	Seasons         []string     `json:"seasons"`
	SourceEpisodeID *uuid.UUID   `json:"sourceEpisodeId"`
}

// Data required to modify the timestamps on a template
type InputTemplateTimestamp struct {
	TemplateID  *uuid.UUID `json:"templateId"`
	TimestampID *uuid.UUID `json:"timestampId"`
}

// Data required to create a new `Timestamp`. See `Timestamp` for a description of each field
type InputTimestamp struct {
	At     float64          `json:"at"`
	TypeID *uuid.UUID       `json:"typeId"`
	Source *TimestampSource `json:"source"`
}

type InputTimestampOn struct {
	// The episode id the timestamp will be created on
	EpisodeID *uuid.UUID `json:"episodeId"`
	// The new values for the timestamp
	Timestamp *InputTimestamp `json:"timestamp"`
}

// Data required to create a new `TimestampType`. See `TimestampType` for a description of each field
type InputTimestampType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// When logging in with a password or refresh token, you can get new tokens and account info
type LoginData struct {
	// A JWT that should be used in the header of all requests: `Authorization: Bearer <authToken>`
	AuthToken string `json:"authToken"`
	// A JWT used for the `loginRefresh` query to get new `LoginData`
	RefreshToken string `json:"refreshToken"`
	// The personal account information of the user that got authenticated
	Account *Account `json:"account"`
}

// Where all the user preferences are stored. This includes what timestamps the user doesn't want to
// watch
type Preferences struct {
	ID        *uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	// The `User.id` that this preferences object belongs to
	UserID *uuid.UUID `json:"userId"`
	// The `User` that the preferences belong to
	User *User `json:"user"`
	// Whether or not the user wants to automatically skip section. Default: `true`
	EnableAutoSkip bool `json:"enableAutoSkip"`
	// Whether or not the user wants to auto-play the videos. Default: `true`
	EnableAutoPlay bool `json:"enableAutoPlay"`
	// Whether or not the bottom toolbar with the video progress and play button is minimized after
	// inactivity while editing
	MinimizeToolbarWhenEditing bool `json:"minimizeToolbarWhenEditing"`
	// When false, timeline is pinned to the bottom of the screen after inactivity. When true, it is
	// hidden completely
	HideTimelineWhenMinimized bool       `json:"hideTimelineWhenMinimized"`
	ColorTheme                ColorTheme `json:"colorTheme"`
	// Whether or not the user whats to skip branding timestamps. Default: `true`
	SkipBranding bool `json:"skipBranding"`
	// Whether or not the user whats to skip regular intros. Default: `true`
	SkipIntros bool `json:"skipIntros"`
	// Whether or not the user whats to skip the first of an intro. Default: `false`
	SkipNewIntros bool `json:"skipNewIntros"`
	// Whether or not the user whats to kip intros that have plot progression rather than the standard animation. Default: `false`
	SkipMixedIntros bool `json:"skipMixedIntros"`
	// Whether or not the user whats to skip recaps at the beginning of episodes. Default: `true`
	SkipRecaps bool `json:"skipRecaps"`
	// Whether or not the user whats to skip filler content. Default: `true`
	SkipFiller bool `json:"skipFiller"`
	// Whether or not the user whats to skip canon content. Default: `false`
	SkipCanon bool `json:"skipCanon"`
	// Whether or not the user whats to skip commertial transitions. Default: `true`
	SkipTransitions bool `json:"skipTransitions"`
	// Whether or not the user whats to skip credits/outros. Default: `true`
	SkipCredits bool `json:"skipCredits"`
	// Whether or not the user whats to skip the first of a credits/outro. Default: `false`
	SkipNewCredits bool `json:"skipNewCredits"`
	// Whether or not the user whats to skip credits/outros that have plot progression rather than the standard animation. Default: `false`
	SkipMixedCredits bool `json:"skipMixedCredits"`
	// Whether or not to skip the next episode's preview. Default: `true`
	SkipPreview bool `json:"skipPreview"`
	// Whether or not to skip an episode's static title card. Default: `true`
	SkipTitleCard bool `json:"skipTitleCard"`
}

// A show containing a list of episodes and relevant links
type Show struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The show name
	//
	// ### Examples
	//
	// - "Death Note"
	// - "My Hero Academia"
	Name string `json:"name"`
	// The show's original Japanese name
	//
	// ### Examples
	//
	// - "Desu Nōto"
	// - "Boku no Hīrō Akademia"
	OriginalName *string `json:"originalName"`
	// A link to the anime's official website
	Website *string `json:"website"`
	// A link to a show poster
	Image *string `json:"image"`
	// The list of admins for the show
	Admins []*ShowAdmin `json:"admins"`
	// All the episodes that belong to the show
	Episodes []*Episode `json:"episodes"`
	// All the templates that belong to this show
	Templates []*Template `json:"templates"`
	// How many seasons are associated with this show
	SeasonCount int `json:"seasonCount"`
	// How many episodes are apart of this show
	EpisodeCount int `json:"episodeCount"`
}

func (Show) IsBaseModel() {}

// A list of users that have elevated permissions when making changes to a show, it's episodes, and
// timestamps. Show admins are responsible for approving any changes that users might submit.
//
// If a user has the `ADMIN` or `DEV` roles, they do not need to be show admins to approve changes or
// make changes directly. Likewise, if a show doesn't have an admin, the user that create the
// show/episode will have temporary access to editing the data until someone becomes that shows admin.
//
// Admins can be created using the API and will soon come to the Anime Skip player/website.
type ShowAdmin struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The `Show.id` that the admin has elevated privileges for
	ShowID *uuid.UUID `json:"showId"`
	// The `Show` that the admin has elevated privileges for
	Show *Show `json:"show"`
	// The `User.id` that the admin privileges belong to
	UserID *uuid.UUID `json:"userId"`
	// The `User` that the admin privileges belong to
	User *User `json:"user"`
}

func (ShowAdmin) IsBaseModel() {}

// When no timestamps exist for a specific episode, templates are setup to provide fallback timestamps
type Template struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The id of the show that this template is for
	ShowID *uuid.UUID `json:"showId"`
	// The show that this template is for
	Show *Show `json:"show"`
	// Specify the scope of the template, if it's for the entire show, or just for a set of seasons
	Type TemplateType `json:"type"`
	// When the template is for a set of seasons, this is the set of seasons it is applied to
	Seasons []string `json:"seasons"`
	// The id of the episode used to create the template. All the timestamps are from this episode
	SourceEpisodeID *uuid.UUID `json:"sourceEpisodeId"`
	// The episode used to create the template. All the timestamps are from this episode
	SourceEpisode *Episode `json:"sourceEpisode"`
	// The list of timestamps that are apart of this template
	Timestamps []*Timestamp `json:"timestamps"`
	// The list of timestamp ids that are apart of this template. Since this is a many-to-many
	// relationship, this field will resolve quicker than `timestamps` since it doesn't have to do an
	// extra join
	//
	// This is useful when you already got the episode and timestamps, and you just need to know what
	// timestamps are apart of the template
	TimestampIds []*uuid.UUID `json:"timestampIds"`
}

func (Template) IsBaseModel() {}

// The many to many object that links a timestamp to a template
type TemplateTimestamp struct {
	TemplateID  *uuid.UUID `json:"templateId"`
	Template    *Template  `json:"template"`
	TimestampID *uuid.UUID `json:"timestampId"`
	Timestamp   *Timestamp `json:"timestamp"`
}

// Episode info provided by a third party. See `Episode` for a description of each field.
//
// When creating data based on this type, fill out and post an episode, then timestamps based on the
// data here. All fields will map 1 to 1 with the exception of `source`. Since a source belongs to a
// episode for third party data, but belongs to timestamps in Anime Skip, the source should be
// propogated down to each of the timestamps. This way when more timestamps are added, a episode can
// have muliple timestamp sources.
//
// > Make sure to fill out the `source` field so that original owner of the timestamp is maintained
type ThirdPartyEpisode struct {
	// The Anime Skip `Episode.id` when the `source` is `ANIME_SKIP`, otherwise this is null
	ID             *uuid.UUID             `json:"id"`
	Season         *string                `json:"season"`
	Number         *string                `json:"number"`
	AbsoluteNumber *string                `json:"absoluteNumber"`
	BaseDuration   *float64               `json:"baseDuration"`
	Name           *string                `json:"name"`
	Source         *TimestampSource       `json:"source"`
	Timestamps     []*ThirdPartyTimestamp `json:"timestamps"`
	// The id of the show from the third party
	ShowID string          `json:"showId"`
	Show   *ThirdPartyShow `json:"show"`
}

type ThirdPartyShow struct {
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ThirdPartyTimestamp struct {
	// The Anime Skip `Timestamp.id` when the `Episode.source` is `ANIME_SKIP`, otherwise this is null
	ID *uuid.UUID `json:"id"`
	// The actual time the timestamp is at
	At float64 `json:"at"`
	// The id specifying the type the timestamp is
	TypeID *uuid.UUID     `json:"typeId"`
	Type   *TimestampType `json:"type"`
}

type Timestamp struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The actual time the timestamp is at
	At     float64         `json:"at"`
	Source TimestampSource `json:"source"`
	// The id specifying the type the timestamp is
	TypeID *uuid.UUID `json:"typeId"`
	// The type the timestamp is. Thid field is a constant string so including it has no effect on
	// performance or query complexity.
	Type *TimestampType `json:"type"`
	// The `Episode.id` that the timestamp belongs to
	EpisodeID *uuid.UUID `json:"episodeId"`
	// The `Episode` that the timestamp belongs to
	Episode *Episode `json:"episode"`
}

func (Timestamp) IsBaseModel() {}

// The type a timestamp can be. This table rarely changes so the values fetched can either be hard
// coded or fetch occasionally. Anime Skip website and web extension use hardcoded maps to store this
// data, but a third party might want to fetch and cache this instead since you won't know when Anime
// Skip adds timestamps
type TimestampType struct {
	ID              *uuid.UUID `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID *uuid.UUID `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID *uuid.UUID `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *uuid.UUID `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	// The name of the timestamp type
	Name string `json:"name"`
	// The description for what this type represents
	Description string `json:"description"`
}

func (TimestampType) IsBaseModel() {}

type UpdatedTimestamps struct {
	Created []*Timestamp `json:"created"`
	Updated []*Timestamp `json:"updated"`
	Deleted []*Timestamp `json:"deleted"`
}

// Information about a user that is public. See `Account` for a description of each field
type User struct {
	ID           *uuid.UUID   `json:"id"`
	CreatedAt    time.Time    `json:"createdAt"`
	DeletedAt    *time.Time   `json:"deletedAt"`
	Username     string       `json:"username"`
	ProfileURL   string       `json:"profileUrl"`
	AdminOfShows []*ShowAdmin `json:"adminOfShows"`
}

// Color theme the user prefers
type ColorTheme string

const (
	// Change to match where you're watching
	ColorThemePerService        ColorTheme = "PER_SERVICE"
	ColorThemeAnimeSkipBlue     ColorTheme = "ANIME_SKIP_BLUE"
	ColorThemeVrvYellow         ColorTheme = "VRV_YELLOW"
	ColorThemeFunimationPurple  ColorTheme = "FUNIMATION_PURPLE"
	ColorThemeCrunchyrollOrange ColorTheme = "CRUNCHYROLL_ORANGE"
)

var AllColorTheme = []ColorTheme{
	ColorThemePerService,
	ColorThemeAnimeSkipBlue,
	ColorThemeVrvYellow,
	ColorThemeFunimationPurple,
	ColorThemeCrunchyrollOrange,
}

func (e ColorTheme) IsValid() bool {
	switch e {
	case ColorThemePerService, ColorThemeAnimeSkipBlue, ColorThemeVrvYellow, ColorThemeFunimationPurple, ColorThemeCrunchyrollOrange:
		return true
	}
	return false
}

func (e ColorTheme) String() string {
	return string(e)
}

func (e *ColorTheme) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ColorTheme(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ColorTheme", str)
	}
	return nil
}

func (e ColorTheme) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Which of the supported services the `EpisodeUrl` was created for. This is a simple enum that allows
// for simple checks, but this data can also be pulled from the url in the case of UNKNOWN
type EpisodeSource string

const (
	// Data came from an external source
	EpisodeSourceUnknown EpisodeSource = "UNKNOWN"
	// Data is from <vrv.co>
	EpisodeSourceVrv EpisodeSource = "VRV"
	// Data is from <funimation.com>
	EpisodeSourceFunimation EpisodeSource = "FUNIMATION"
	// Data is from <crunchyroll.com> and <beta.crunchyroll.com>
	EpisodeSourceCrunchyroll EpisodeSource = "CRUNCHYROLL"
)

var AllEpisodeSource = []EpisodeSource{
	EpisodeSourceUnknown,
	EpisodeSourceVrv,
	EpisodeSourceFunimation,
	EpisodeSourceCrunchyroll,
}

func (e EpisodeSource) IsValid() bool {
	switch e {
	case EpisodeSourceUnknown, EpisodeSourceVrv, EpisodeSourceFunimation, EpisodeSourceCrunchyroll:
		return true
	}
	return false
}

func (e EpisodeSource) String() string {
	return string(e)
}

func (e *EpisodeSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EpisodeSource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EpisodeSource", str)
	}
	return nil
}

func (e EpisodeSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A user's role in the system. Higher roles allow a user write access to certain data that a normal
// user would not. Some queries and mutations are only alloed by certain roles
type Role string

const (
	// Highest role. Has super user access to all queries and mutations
	RoleDev Role = "DEV"
	// Administrator role. Has some elevated permissions
	RoleAdmin Role = "ADMIN"
	// Basic role. Has no elevated permissions
	RoleUser Role = "USER"
)

var AllRole = []Role{
	RoleDev,
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleDev, RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The scope that a template applies to
type TemplateType string

const (
	// The template is loaded for all episodes of a given show
	TemplateTypeShow TemplateType = "SHOW"
	// The template is loaded for episodes of a given show where their season is included in `Template.seasons`
	TemplateTypeSeasons TemplateType = "SEASONS"
)

var AllTemplateType = []TemplateType{
	TemplateTypeShow,
	TemplateTypeSeasons,
}

func (e TemplateType) IsValid() bool {
	switch e {
	case TemplateTypeShow, TemplateTypeSeasons:
		return true
	}
	return false
}

func (e TemplateType) String() string {
	return string(e)
}

func (e *TemplateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TemplateType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TemplateType", str)
	}
	return nil
}

func (e TemplateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Where a timestamp originated from
type TimestampSource string

const (
	TimestampSourceAnimeSkip TimestampSource = "ANIME_SKIP"
	TimestampSourceBetterVrv TimestampSource = "BETTER_VRV"
)

var AllTimestampSource = []TimestampSource{
	TimestampSourceAnimeSkip,
	TimestampSourceBetterVrv,
}

func (e TimestampSource) IsValid() bool {
	switch e {
	case TimestampSourceAnimeSkip, TimestampSourceBetterVrv:
		return true
	}
	return false
}

func (e TimestampSource) String() string {
	return string(e)
}

func (e *TimestampSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimestampSource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TimestampSource", str)
	}
	return nil
}

func (e TimestampSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
