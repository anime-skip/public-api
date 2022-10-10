package main

import (
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/alerts"
	"anime-skip.com/public-api/internal/config"
	"anime-skip.com/public-api/internal/graphql/handler"
	"anime-skip.com/public-api/internal/http"
	"anime-skip.com/public-api/internal/jwt"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres"
	"anime-skip.com/public-api/internal/utils"
	"github.com/samber/lo"
)

// Compile time constants
var (
	VERSION string
	STAGE   string
)

func main() {
	log.I("Starting anime-skip/public-api")

	pg := postgres.Open(
		config.DatabaseURL(),
		config.DatabaseDisableSSL(),
		lo.ToPtr(config.DatabaseVersion()),
		config.DatabaseEnableSeeding(),
	)

	anilist := http.NewAnilistService()

	pgEpisodeService := postgres.NewEpisodeService(pg)
	pgEpisodeURLService := postgres.NewEpisodeURLService(pg)
	pgShowAdminService := postgres.NewShowAdminService(pg)
	pgTemplateService := postgres.NewTemplateService(pg)
	pgUserService := postgres.NewUserService(pg)

	jwtAuthService := jwt.NewJWTAuthService(
		config.JWTSecret(),
		pgUserService,
	)
	httpEmailService := http.NewAnimeSkipEmailService(
		config.EmailServiceHost(),
		config.EmailServiceSecret(),
		config.EmailServiceEnabled(),
	)
	httpRecaptchaService := http.NewGoogleRecaptchaService(
		config.RecaptchaSecret(),
		config.RecaptchaResponseAllowList(),
	)
	betterVRV := http.NewBetterVRVThirdPartyService(
		config.BetterVRVAppID(),
		config.BetterVRVAPIKey(),
	)

	rateLimiter := http.NewRequestRateLimiter()
	discord, err := alerts.NewDiscordAPIClient(config.DiscordBotToken(), config.DiscordAlertChannelID())
	if err != nil {
		panic(err)
	}

	services := internal.Services{
		APIClientService:         postgres.NewAPIClientService(pg),
		AuthService:              jwtAuthService,
		EmailService:             httpEmailService,
		EpisodeService:           pgEpisodeService,
		EpisodeURLService:        pgEpisodeURLService,
		ExternalLinkService:      postgres.NewExternalLinkService(pg),
		PreferencesService:       postgres.NewPreferencesService(pg),
		RecaptchaService:         httpRecaptchaService,
		ShowAdminService:         pgShowAdminService,
		ShowService:              postgres.NewShowService(pg, anilist),
		TemplateService:          pgTemplateService,
		TemplateTimestampService: postgres.NewTemplateTimestampService(pg),
		TimestampService:         postgres.NewTimestampService(pg),
		TimestampTypeService:     postgres.NewTimestampTypeService(pg),
		UserService:              pgUserService,
		UserReportService:        postgres.NewUserReportService(pg, discord),
		ThirdPartyService: utils.NewAggregateThirdPartyService([]internal.ThirdPartyService{
			postgres.NewThirdPartyService(pg),
			utils.NewCachedThirdPartyService(betterVRV, 30*time.Minute),
		}),
	}

	graphqlHandler := handler.NewGraphqlHandler(
		pg,
		services,
		rateLimiter,
		config.EnableIntrospection(),
	)

	server := http.NewChiServer(
		config.Port(),
		config.EnablePlayground(),
		"/graphql",
		graphqlHandler,
		services,
		rateLimiter,
		VERSION,
		STAGE,
		internal.SHARED_CLIENT_ID,
	)

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
