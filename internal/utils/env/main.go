package env

import (
	"os"
	"strconv"
	"strings"

	log "anime-skip.com/backend/internal/utils/log"
)

// General
var IS_DEV bool
var IS_STAGED bool
var GIN_MODE string
var PORT int
var LOG_LEVEL int
var JWT_SECRET string

// Feature Flags
var LOG_SQL bool
var ENABLE_COLOR_LOGS bool
var ENABLE_INTROSPECTION bool
var ENABLE_PLAYGROUND bool
var DISABLE_SHOW_ADMIN_DIRECTIVE bool
var DISABLE_EMAILS bool
var SLEEP_BAN_IP bool
var BANNED_IP_ADDRESSES []string
var DISABLE_RATE_LIMITTING bool

// Database
var DATABASE_URL string
var DATABASE_DISABLE_SSL bool
var DATABASE_MIGRATION *int
var DATABASE_ENABLE_SEEDING bool
var DATABASE_SKIP_MIGRATIONS bool

// Emails
var EMAIL_SERVICE_HOST string
var EMAIL_SERVICE_SECRET string

// 3rd Party Services
var RECAPTCHA_SECRET string
var RECAPTCHA_RESPONSE_ALLOWLIST []string
var BETTER_VRV_APP_ID string
var BETTER_VRV_API_KEY string

func init() {
	IS_DEV = envBoolOrFalse("IS_DEV")
	IS_STAGED = envBoolOrFalse("IS_STAGED")
	GIN_MODE = envStringOr("GIN_MODE", "development")
	PORT = envIntOr("PORT", 8081)
	LOG_LEVEL = envIntOr("LOG_LEVEL", 0)
	JWT_SECRET = envString("JWT_SECRET")
	BANNED_IP_ADDRESSES = envStringArray("BANNED_IP_ADDRESSES")
	SLEEP_BAN_IP = envBoolOrFalse("SLEEP_BAN_IP")
	DISABLE_RATE_LIMITTING = envBoolOrFalse("DISABLE_RATE_LIMITTING")
	LOG_SQL = envBoolOrFalse("LOG_SQL")
	ENABLE_COLOR_LOGS = envBoolOrFalse("ENABLE_COLOR_LOGS")
	ENABLE_INTROSPECTION = envBoolOrFalse("ENABLE_INTROSPECTION")
	ENABLE_PLAYGROUND = envBoolOrFalse("ENABLE_PLAYGROUND")
	DISABLE_SHOW_ADMIN_DIRECTIVE = envBoolOrFalse("DISABLE_SHOW_ADMIN_DIRECTIVE")
	DISABLE_EMAILS = envBoolOrFalse("DISABLE_EMAILS")
	DATABASE_URL = envString("DATABASE_URL")
	DATABASE_DISABLE_SSL = envBoolOrFalse("DATABASE_DISABLE_SSL")
	DATABASE_MIGRATION = envIntOrNil("DATABASE_MIGRATION")
	DATABASE_ENABLE_SEEDING = envBoolOrFalse("DATABASE_ENABLE_SEEDING")
	DATABASE_SKIP_MIGRATIONS = envBoolOrFalse("DATABASE_SKIP_MIGRATIONS")
	EMAIL_SERVICE_HOST = envString("EMAIL_SERVICE_HOST")
	EMAIL_SERVICE_SECRET = envString("EMAIL_SERVICE_SECRET")
	RECAPTCHA_SECRET = envString("RECAPTCHA_SECRET")
	RECAPTCHA_RESPONSE_ALLOWLIST = envStringArray("RECAPTCHA_RESPONSE_ALLOWLIST")
	BETTER_VRV_APP_ID = envString("BETTER_VRV_APP_ID")
	BETTER_VRV_API_KEY = envString("BETTER_VRV_API_KEY")

	log.I("Loaded ENV Variables")
	log.V("IS_DEV=%v", IS_DEV)
	log.V("IS_STAGED=%v", IS_STAGED)
	log.V("GIN_MODE=%v", GIN_MODE)
	log.V("PORT=%v", PORT)
	log.V("LOG_LEVEL=%v", LOG_LEVEL)
	log.V("LOG_SQL=%v", LOG_SQL)
	log.V("BANNED_IP_ADDRESSES=%v", BANNED_IP_ADDRESSES)
	log.V("ENABLE_COLOR_LOGS=%v", ENABLE_COLOR_LOGS)
	log.V("ENABLE_INTROSPECTION=%v", ENABLE_INTROSPECTION)
	log.V("ENABLE_PLAYGROUND=%v", ENABLE_PLAYGROUND)
	log.V("DISABLE_RATE_LIMITTING=%v", DISABLE_RATE_LIMITTING)
	log.V("DISABLE_SHOW_ADMIN_DIRECTIVE=%v", DISABLE_SHOW_ADMIN_DIRECTIVE)
	log.V("DISABLE_EMAILS=%v", DISABLE_EMAILS)
	log.V("DATABASE_MIGRATION=%v", DATABASE_MIGRATION)
	log.V("DATABASE_ENABLE_SEEDING=%v", DATABASE_ENABLE_SEEDING)
}

func envString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return strings.TrimSpace(v)
}

func envStringOr(k, defaultValue string) string {
	v := os.Getenv(k)
	if v == "" {
		log.V("ENV missing (%s), defaulting to %v", k, defaultValue)
		return defaultValue
	}
	return strings.TrimSpace(v)
}

func envStringArray(k string) []string {
	str := envStringOr(k, "")
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}

func envBoolOrFalse(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.W("%s missing from ENV, defaulting to false", k)
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panic("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}

func envIntOr(k string, defaultValue int) int {
	i, err := strconv.Atoi(os.Getenv(k))
	if err != nil {
		return defaultValue
	}
	return i
}

func envIntOrNil(k string) *int {
	str := os.Getenv(k)
	if str == "" {
		return nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("ENV err: [" + k + "]\n" + err.Error())
	}
	return &i
}
