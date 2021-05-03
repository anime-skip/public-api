package health_service

import (
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/env"
)

type ServerConfig struct {
	Version        string  `json:"version"`
	Playground     bool    `json:"playground"`
	PlaygroundPath *string `json:"playgroundPath,omitempty"`
	Introspection  bool    `json:"introspection"`
}

func GetServerConfig() ServerConfig {
	isPlaygroundEnabled := env.ENABLE_PLAYGROUND
	version := constants.VERSION
	if env.IS_DEV || env.IS_STAGED {
		version += constants.VERSION_SUFFIX
	}

	serverConfig := ServerConfig{
		Version:       version,
		Playground:    isPlaygroundEnabled,
		Introspection: env.ENABLE_INTROSPECTION,
	}
	if isPlaygroundEnabled {
		path := "/graphiql"
		serverConfig.PlaygroundPath = &path
	}

	return serverConfig
}
