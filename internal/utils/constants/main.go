package constants

const (
	CTX_GIN_CONTEXT  = "ctx_gin_context"
	CTX_USER_ID      = "ctx_user_id"
	CTX_JWT_ERROR    = "ctx_jwt_error"
	CTX_ROLE         = "ctx_role"
	CTX_DATA_LOADERS = "ctx_data_loaders"
	CTX_API_CLIENT   = "ctx_api_client"
)

const (
	ROLE_DEV   = 0
	ROLE_ADMIN = 1
	ROLE_USER  = 2
)

const (
	EPISODE_SOURCE_UNKNOWN    = 0
	EPISODE_SOURCE_VRV        = 1
	EPISODE_SOURCE_FUNIMATION = 2
)

const (
	TIMESTAMP_SOURCE_ANIME_SKIP = 0
	TIMESTAMP_SOURCE_BETTER_VRV = 1
)

const (
	TEMPLATE_TYPE_SHOW    = 0
	TEMPLATE_TYPE_SEASONS = 1
)

const (
	LOG_LEVEL_VERBOSE = 0
	LOG_LEVEL_DEBUG   = 1
	LOG_LEVEL_WARNING = 2
	LOG_LEVEL_ERROR   = 3
)

const (
	TIMESTAMP_ID_CANON         = "9edc0037-fa4e-47a7-a29a-d9c43368daa8"
	TIMESTAMP_ID_MUST_WATCH    = "e384759b-3cd2-4824-9569-128363b4452b"
	TIMESTAMP_ID_BRANDING      = "97e3629a-95e5-4b1a-9411-73a47c0d0e25"
	TIMESTAMP_ID_INTRO         = "14550023-2589-46f0-bfb4-152976506b4c"
	TIMESTAMP_ID_INTRO_INTRO   = "cbb42238-d285-4c88-9e91-feab4bb8ae0a"
	TIMESTAMP_ID_NEW_INTRO     = "679fb610-ff3c-4cf4-83c0-75bcc7fe8778"
	TIMESTAMP_ID_RECAP         = "f38ac196-0d49-40a9-8fcf-f3ef2f40f127"
	TIMESTAMP_ID_FILLER        = "c48f1dce-1890-4394-8ce6-c3f5b2f95e5e"
	TIMESTAMP_ID_TRANSITION    = "9f0c6532-ccae-4238-83ec-a2804fe5f7b0"
	TIMESTAMP_ID_CREDITS       = "2a730a51-a601-439b-bc1f-7b94a640ffb9"
	TIMESTAMP_ID_MIXED_CREDITS = "6c4ade53-4fee-447f-89e4-3bb29184e87a"
	TIMESTAMP_ID_NEW_CREDITS   = "d839cdb1-21b3-455d-9c21-7ffeb37adbec"
	TIMESTAMP_ID_PREVIEW       = "c7b1eddb-defa-4bc6-a598-f143081cfe4b"
	TIMESTAMP_ID_TITLE_CARD    = "67321535-a4ea-4f21-8bed-fb3c8286b510"
	TIMESTAMP_ID_UNKNOWN       = "ae57fcf9-27b0-49a7-9a99-a91aa7518a29"
)

const (
	LOGIN_RETRY_INCREMENT = 500  // ms
	LOGIN_RETRY_FREEBEES  = 5    // attempts
	LOGIN_RETRY_MAX_SLEEP = 5000 // ms
)
