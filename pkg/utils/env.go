package utils

import (
	"log"
	"os"
	"strconv"
)

// EnvString will return the env or panic if it is not present
func EnvString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	return v
}

// EnvBool will return the env as boolean or panic if it is not present
func EnvBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicln("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}
