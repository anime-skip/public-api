package main

import (
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/http"
	"anime-skip.com/timestamps-service/internal/jwt"
	"anime-skip.com/timestamps-service/internal/postgres"
)

func main() {
	db := postgres.Open(
		config.RequireEnvString("DATABASE_URL"),
		config.EnvBool("DATABASE_DISABLE_SSL"),
	)

	graphqlHandler := graphql.NewGraphqlHandler(
		db,
		config.EnvBool("ENABLE_INTROSPECTION"),
	)

	authenticator := jwt.NewJWTAuthenticator()
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
