package config

import (
	"os"
	"strconv"
	"strings"
)

func envString(key string) string {
	return os.Getenv(key)
}

func envStringOr(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.TrimSpace(value)
}

func requireEnvString(key string) string {
	value := envString(key)
	if strings.TrimSpace(value) == "" {
		panic(key + " is not an environment variable")
	}
	return value
}

func envStringArray(key string) []string {
	str := envStringOr(key, "")
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}

func envBool(key string) bool {
	return envString(key) == "true"
}

func envInt(key string) int {
	value, err := strconv.Atoi(requireEnvString(key))
	if err != nil {
		panic(err)
	}
	return value
}

func envIntOr(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return envInt(key)
}
