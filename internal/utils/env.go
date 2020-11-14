package utils

import (
	"os"
	"strconv"
	"strings"

	log "anime-skip.com/backend/internal/utils/log"
)

// EnvString will return the env as a string or crash if it doesn't exist
func EnvString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// EnvString will return the env as a []string or default to []
func EnvStringArray(k string) []string {
	str := EnvString(k)
	return strings.Split(str, ",")
}

// EnvBool will return the env as a boolean, return false when not set, and panic if it's not a boolean
func EnvBool(k string) bool {
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

// EnvIntOrDefault will return the env as an int or the default value if it doesn't exist
func EnvIntOrDefault(k string, defaultValue int) int {
	i, err := strconv.Atoi(os.Getenv(k))
	if err != nil {
		return defaultValue
	}
	return i
}
