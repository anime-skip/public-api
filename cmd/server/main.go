package main

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/graphql/handler"
	"anime-skip.com/timestamps-service/internal/http"
	"anime-skip.com/timestamps-service/internal/postgres"
)

func main() {
	db := postgres.Open(
		config.RequireEnvString("DATABASE_URL"),
		config.EnvBool("DATABASE_DISABLE_SSL"),
		config.EnvInt("DATABASE_VERSION"),
	)

	services := internal.Services{
		UserService: postgres.NewUserService(db),
	}

	graphqlHandler := handler.NewGraphqlHandler(
		services,
		config.EnvBool("ENABLE_INTROSPECTION"),
	)

	authenticator := http.NewNoAuthenticator()
	server := http.NewChiServer(
		config.EnvInt("PORT"),
		config.EnvBool("ENABLE_PLAYGROUND"),
		"/graphql",
		graphqlHandler,
		authenticator,
	)

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
