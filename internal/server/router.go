package server

import (
	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server/handlers"
	"anime-skip.com/backend/internal/server/middleware"
	"anime-skip.com/backend/internal/utils/env"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

const (
	STATUS_ENDPOINT     = "/status"
	GRAPHQL_ENDPOINT    = "/graphql"
	PLAYGROUND_ENDPOINT = "/graphiql"
)

func GetRoutes(orm *database.ORM) *chi.Mux {
	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(middleware.IPInContext)
	if env.IS_DEV {
		router.Use(chiMiddleware.Logger)
	}
	router.Use(chiMiddleware.Recoverer)

	// Middleware
	router.Use(middleware.Authorization)
	router.Use(middleware.Cors)

	// REST endpoints
	router.Get(STATUS_ENDPOINT, handlers.Status)

	// GraphQL endpoints
	if env.ENABLE_PLAYGROUND {
		router.Get(PLAYGROUND_ENDPOINT, handlers.PlaygroundHandler(GRAPHQL_ENDPOINT))
	}
	router.Post(GRAPHQL_ENDPOINT, handlers.GraphQLHandler(orm))

	return router
}
