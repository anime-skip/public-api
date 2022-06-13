package main

import (
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/config"
	"anime-skip.com/public-api/internal/graphql/handler"
	"anime-skip.com/public-api/internal/http"
	"anime-skip.com/public-api/internal/jwt"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres"
	"anime-skip.com/public-api/internal/utils"
)

// Compile time constants
var (
	VERSION string
	STAGE   string
)

func main() {
	log.I("Starting anime-skip/public-api")

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
	betterVRV := http.NewBetterVRVThirdPartyService(
		config.EnvString("BETTER_VRV_APP_ID"),
		config.EnvString("BETTER_VRV_API_KEY"),
	)

	rateLimiter := http.NewRequestRateLimiter()

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
		ThirdPartyService: utils.NewAggregateThirdPartyService([]internal.ThirdPartyService{
			postgres.NewThirdPartyService(pg),
			utils.NewCachedThirdPartyService(betterVRV, 30*time.Minute),
		}),
	}

	graphqlHandler := handler.NewGraphqlHandler(
		pg,
		services,
		rateLimiter,
		config.EnvBool("ENABLE_INTROSPECTION"),
	)

	server := http.NewChiServer(
		config.EnvInt("PORT"),
		config.EnvBool("ENABLE_PLAYGROUND"),
		"/graphql",
		graphqlHandler,
		services,
		rateLimiter,
		VERSION,
		STAGE,
		internal.SHARED_CLIENT_ID,
	)

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
