package main

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/graphql/handler"
	"anime-skip.com/timestamps-service/internal/http"
	"anime-skip.com/timestamps-service/internal/jwt"
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

	authService := jwt.NewJWTAuthService(
		config.RequireEnvString("JWT_SECRET"),
	)
	emailService := http.NewAnimeSkipEmailService(
		config.RequireEnvString("EMAIL_SERVICE_HOST"),
		config.RequireEnvString("EMAIL_SERVICE_SECRET"),
		config.EnvBool("EMAIL_SERVICE_ENABLED"),
	)
	recaptchaService := http.NewGoogleRecaptchaService(
		config.RequireEnvString("RECAPTCHA_SECRET"),
		config.EnvStringArray("RECAPTCHA_RESPONSE_ALLOWLIST"),
	)
	services := internal.Services{
		APIClientService:         postgres.NewAPIClientService(db),
		AuthService:              authService,
		EmailService:             emailService,
		EpisodeService:           postgres.NewEpisodeService(db),
		EpisodeURLService:        postgres.NewEpisodeURLService(db),
		PreferencesService:       postgres.NewPreferencesService(db),
		RecaptchaService:         recaptchaService,
		ShowAdminService:         postgres.NewShowAdminService(db),
		ShowService:              postgres.NewShowService(db),
		TemplateService:          postgres.NewTemplateService(db),
		TemplateTimestampService: postgres.NewTemplateTimestampService(db),
		TimestampService:         postgres.NewTimestampService(db),
		TimestampTypeService:     postgres.NewTimestampTypeService(db),
		UserService:              postgres.NewUserService(db),
	}

	graphqlHandler := handler.NewGraphqlHandler(
		db,
		services,
		config.EnvBool("ENABLE_INTROSPECTION"),
	)

	server := http.NewChiServer(
		config.EnvInt("PORT"),
		config.EnvBool("ENABLE_PLAYGROUND"),
		"/graphql",
		graphqlHandler,
		authService,
	)

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
