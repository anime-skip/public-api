package utils

import (
	"os"
	"strconv"
	"strings"

	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

// EnvString will return the env as a boolean or default to false
func EnvString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// EnvString will return the env as a boolean or default to false
func EnvStringArray(k string) []string {
	str := EnvString(k)
	return strings.Split(str, ",")
}

// EnvBool will return the env as boolean or panic if it is not present
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
