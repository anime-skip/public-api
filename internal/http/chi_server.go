package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"anime-skip.com/public-api/internal/graphql/playground"
	"github.com/go-chi/chi"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
)

type chiServer struct {
	port               int
	enablePlayground   bool
	graphqlPath        string
	graphqlHandler     internal.GraphQLHandler
	services           internal.Services
	version            string
	stage              string
	playgroundClientID string
	rateLimiter        internal.RateLimiter
}

func NewChiServer(
	port int,
	enablePlayground bool,
	graphqlPath string,
	graphqlHandler internal.GraphQLHandler,
	services internal.Services,
	rateLimiter internal.RateLimiter,
	version string,
	stage string,
	playgroundClientID string,
) internal.Server {
	log.D("Using Chi for routing...")
	return &chiServer{
		port:               port,
		enablePlayground:   enablePlayground,
		graphqlPath:        graphqlPath,
		graphqlHandler:     graphqlHandler,
		services:           services,
		version:            version,
		stage:              stage,
		playgroundClientID: playgroundClientID,
		rateLimiter:        rateLimiter,
	}
}

type Middleware = func(next http.Handler) http.Handler

func (s *chiServer) Start() error {
	r := chi.NewRouter()
	r.Use(s.corsMiddleware)
	r.Use(s.ipMiddleware)

	r.Get("/status", s.statusHandler)
	if s.enablePlayground {
		r.Handle("/", playground.Handler("Anime Skip Playground", s.graphqlPath, s.playgroundClientID))
	}
	r.Route(s.graphqlPath, func(r chi.Router) {
		r.Use(s.directivesMiddleware)
		r.Use(s.clientIDMiddleware)
		if s.rateLimiter != nil {
			r.Use(s.rateLimiter.HttpMiddleware)
		}
		r.Handle("/", s.graphqlHandler.Handler)
	})

	log.I("Started server @ :%d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
}

func (s *chiServer) statusHandler(rw http.ResponseWriter, r *http.Request) {
	status := internal.ApiStatus{
		Version:       s.version,
		Stage:         s.stage,
		Status:        "RUNNING",
		Playground:    s.enablePlayground,
		Introspection: s.graphqlHandler.EnableIntrospection,
	}
	writeJson(rw, status, http.StatusOK)
}

func (s *chiServer) directivesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithServices(ctx, s.services)
		ctx = context.WithAuthToken(ctx, getAuthToken(r))
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (s *chiServer) clientIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := utils.Ptr(r.Header.Get("X-Client-ID"))
		if clientID == nil || strings.TrimSpace(*clientID) == "" {
			writeGraphqlError(w, "The X-Client-ID header must be passed", http.StatusForbidden)
			return
		}

		client, err := s.services.APIClientService.Get(r.Context(), internal.APIClientsFilter{
			ID: utils.Ptr(strings.TrimSpace(*clientID)),
		})
		if err != nil {
			writeGraphqlError(w, "Invalid X-Client-ID header, API client not found", http.StatusForbidden)
			return
		}

		ctx := context.WithAPIClient(r.Context(), client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *chiServer) ipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-FORWARDED-FOR")
		if ip == "" {
			ip = r.RemoteAddr
		}
		ctx := r.Context()
		ctx = context.WithIPAddress(ctx, ip)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (s *chiServer) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Client-ID")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJson(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func writeGraphqlError(w http.ResponseWriter, message string, status int) {
	writeJson(w, map[string]any{
		"errors": []map[string]any{{
			"message": message,
		}},
	}, status)
}

func getAuthToken(r *http.Request) string {
	header := r.Header.Get("authorization")
	re := regexp.MustCompile(`Bearer (.*?\..*?\..*)`)
	matches := re.FindStringSubmatch(header)
	if len(matches) == 0 {
		return ""
	}
	return matches[1]
}
