package main

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/graphql/handler"
	"anime-skip.com/timestamps-service/internal/http"
	"anime-skip.com/timestamps-service/internal/log"
	"anime-skip.com/timestamps-service/internal/postgres"
)

func main() {
	log.I("Starting anime-skip/timestamps-service")

	db := postgres.Open(
		config.RequireEnvString("DATABASE_URL"),
		config.EnvBool("DATABASE_DISABLE_SSL"),
		config.EnvInt("DATABASE_VERSION"),
	)

	services := internal.Services{
		APIClientService:         postgres.NewAPIClientService(db),
		EpisodeService:           postgres.NewEpisodeService(db),
		EpisodeURLService:        postgres.NewEpisodeURLService(db),
		PreferencesService:       postgres.NewPreferencesService(db),
		ShowAdminService:         postgres.NewShowAdminService(db),
		ShowService:              postgres.NewShowService(db),
		TemplateService:          postgres.NewTemplateService(db),
		TemplateTimestampService: postgres.NewTemplateTimestampService(db),
		TimestampService:         postgres.NewTimestampService(db),
		TimestampTypeService:     postgres.NewTimestampTypeService(db),
		UserService:              postgres.NewUserService(db),
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
