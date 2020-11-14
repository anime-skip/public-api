package utils

import (
	"os"
	"strconv"
	"strings"

	log "anime-skip.com/backend/internal/utils/log"
)

type env struct {
	// General
	IS_DEV     bool
	GIN_MODE   string
	PORT       int
	LOG_LEVEL  int
	JWT_SECRET string

	// Feature Flags
	LOG_SQL                      bool
	ENABLE_COLOR_LOGS            bool
	ENABLE_INTROSPECTION         bool
	ENABLE_PLAYGROUND            bool
	DISABLE_SHOW_ADMIN_DIRECTIVE bool
	DISABLE_EMAILS               bool

	// Database
	DATABASE_URL            string
	DATABASE_DISABLE_SSL    bool
	DATABASE_MIGRATION      *int
	DATABASE_ENABLE_SEEDING bool

	// Emails
	EMAIL_SERVICE_HOST   string
	EMAIL_SERVICE_SECRET string

	// 3rd Party Services
	RECAPTCHA_SECRET             string
	RECAPTCHA_RESPONSE_ALLOWLIST []string
	BETTER_VRV_APP_ID            string
	BETTER_VRV_API_KEY           string
}

var ENV env

func init() {
	ENV = env{
		IS_DEV:                       envBoolOrFalse("IS_DEV"),
		GIN_MODE:                     envStringOr("GIN_MODE", "development"),
		PORT:                         envIntOr("PORT", 8081),
		LOG_LEVEL:                    envIntOr("LOG_LEVEL", 0),
		JWT_SECRET:                   envString("JWT_SECRET"),
		LOG_SQL:                      envBoolOrFalse("LOG_SQL"),
		ENABLE_COLOR_LOGS:            envBoolOrFalse("ENABLE_COLOR_LOGS"),
		ENABLE_INTROSPECTION:         envBoolOrFalse("ENABLE_INTROSPECTION"),
		ENABLE_PLAYGROUND:            envBoolOrFalse("ENABLE_PLAYGROUND"),
		DISABLE_SHOW_ADMIN_DIRECTIVE: envBoolOrFalse("DISABLE_SHOW_ADMIN_DIRECTIVE"),
		DISABLE_EMAILS:               envBoolOrFalse("DISABLE_EMAILS"),
		DATABASE_URL:                 envString("DATABASE_URL"),
		DATABASE_DISABLE_SSL:         envBoolOrFalse("DATABASE_DISABLE_SSL"),
		DATABASE_MIGRATION:           envIntOrNil("DATABASE_MIGRATION"),
		DATABASE_ENABLE_SEEDING:      envBoolOrFalse("DATABASE_ENABLE_SEEDING"),
		EMAIL_SERVICE_HOST:           envString("EMAIL_SERVICE_HOST"),
		EMAIL_SERVICE_SECRET:         envString("EMAIL_SERVICE_SECRET"),
		RECAPTCHA_SECRET:             envString("RECAPTCHA_SECRET"),
		RECAPTCHA_RESPONSE_ALLOWLIST: envStringArray("RECAPTCHA_RESPONSE_ALLOWLIST"),
		BETTER_VRV_APP_ID:            envString("BETTER_VRV_APP_ID"),
		BETTER_VRV_API_KEY:           envString("BETTER_VRV_API_KEY"),
	}
	log.I("Loaded ENV Variables")
	log.V("IS_DEV=%v", ENV.IS_DEV)
	log.V("GIN_MODE=%v", ENV.GIN_MODE)
	log.V("PORT=%v", ENV.PORT)
	log.V("LOG_LEVEL=%v", ENV.LOG_LEVEL)
	log.V("LOG_SQL=%v", ENV.LOG_SQL)
	log.V("ENABLE_COLOR_LOGS=%v", ENV.ENABLE_COLOR_LOGS)
	log.V("ENABLE_INTROSPECTION=%v", ENV.ENABLE_INTROSPECTION)
	log.V("ENABLE_PLAYGROUND=%v", ENV.ENABLE_PLAYGROUND)
	log.V("DISABLE_SHOW_ADMIN_DIRECTIVE=%v", ENV.DISABLE_SHOW_ADMIN_DIRECTIVE)
	log.V("DISABLE_EMAILS=%v", ENV.DISABLE_EMAILS)
	log.V("DATABASE_MIGRATION=%v", ENV.DATABASE_MIGRATION)
	log.V("DATABASE_ENABLE_SEEDING=%v", ENV.DATABASE_ENABLE_SEEDING)
}

func envString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

func envStringOr(k, defaultValue string) string {
	v := os.Getenv(k)
	if v == "" {
		log.V("ENV missing (), defaulting to %", k, defaultValue)
		return defaultValue
	}
	return v
}

func envStringArray(k string) []string {
	str := envString(k)
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
