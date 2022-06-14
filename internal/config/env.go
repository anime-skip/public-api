package config

import (
	"os"
	"strconv"
	"strings"
)

func EnvString(key string) string {
	return os.Getenv(key)
}

func EnvStringOr(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.TrimSpace(value)
}

func RequireEnvString(key string) string {
	value := EnvString(key)
	if strings.TrimSpace(value) == "" {
		panic(key + " is not an environment variable")
	}
	return value
}

func EnvStringArray(key string) []string {
	str := EnvStringOr(key, "")
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}

func EnvBool(key string) bool {
	return EnvString(key) == "true"
}

func EnvInt(key string) int {
	value, err := strconv.Atoi(RequireEnvString(key))
	if err != nil {
		panic(err)
	}
	return value
}

func EnvIntOr(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return EnvInt(key)
}
