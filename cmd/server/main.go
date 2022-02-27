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

	pg := postgres.Open(
		config.RequireEnvString("DATABASE_URL"),
		config.EnvBool("DATABASE_DISABLE_SSL"),
		config.EnvInt("DATABASE_VERSION"),
		config.EnvBool("DATABASE_ENABLE_SEEDING"),
	)

	pgEpisodeService := postgres.NewEpisodeService(pg)
	pgEpisodeURLService := postgres.NewEpisodeURLService(pg)
	pgShowAdminService := postgres.NewShowAdminService(pg)
	pgTemplateService := postgres.NewTemplateService(pg)

	jwtAuthService := jwt.NewJWTAuthService(
		config.RequireEnvString("JWT_SECRET"),
	)
	httpEmailService := http.NewAnimeSkipEmailService(
		config.RequireEnvString("EMAIL_SERVICE_HOST"),
		config.RequireEnvString("EMAIL_SERVICE_SECRET"),
		config.EnvBool("EMAIL_SERVICE_ENABLED"),
	)
	httpRecaptchaService := http.NewGoogleRecaptchaService(
		config.RequireEnvString("RECAPTCHA_SECRET"),
		config.EnvStringArray("RECAPTCHA_RESPONSE_ALLOWLIST"),
	)

	services := internal.Services{
		APIClientService:         postgres.NewAPIClientService(pg),
		AuthService:              jwtAuthService,
		EmailService:             httpEmailService,
		EpisodeService:           pgEpisodeService,
		EpisodeURLService:        pgEpisodeURLService,
		PreferencesService:       postgres.NewPreferencesService(pg),
		RecaptchaService:         httpRecaptchaService,
		ShowAdminService:         pgShowAdminService,
		ShowService:              postgres.NewShowService(pg),
		TemplateService:          pgTemplateService,
		TemplateTimestampService: postgres.NewTemplateTimestampService(pg),
		TimestampService:         postgres.NewTimestampService(pg),
		TimestampTypeService:     postgres.NewTimestampTypeService(pg),
		UserService:              postgres.NewUserService(pg),
	}
	directiveServices := internal.DirectiveServices{
		AuthService:       jwtAuthService,
		EpisodeService:    pgEpisodeService,
		EpisodeURLService: pgEpisodeURLService,
		ShowAdminService:  pgShowAdminService,
		TemplateService:   pgTemplateService,
	}

	graphqlHandler := handler.NewGraphqlHandler(
		pg,
		services,
		config.EnvBool("ENABLE_INTROSPECTION"),
	)

	server := http.NewChiServer(
		config.EnvInt("PORT"),
		config.EnvBool("ENABLE_PLAYGROUND"),
		"/graphql",
		graphqlHandler,
		directiveServices,
	)

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
