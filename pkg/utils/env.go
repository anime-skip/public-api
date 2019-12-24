package utils

import (
	"os"
	"strconv"

	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

// EnvString will return the env or panic if it is not present
func EnvString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// EnvBool will return the env as boolean or panic if it is not present
func EnvBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.D("%s missing from ENV, defaulting to false", k)
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panic("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}
