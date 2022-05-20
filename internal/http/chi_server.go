package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
)

type chiServer struct {
	port             int
	enablePlayground bool
	graphqlPath      string
	graphqlHandler   internal.GraphQLHandler
	services         internal.Services
	version          string
	stage            string
}

func NewChiServer(
	port int,
	enablePlayground bool,
	graphqlPath string,
	graphqlHandler internal.GraphQLHandler,
	services internal.Services,
	version string,
	stage string,
) internal.Server {
	log.D("Using Chi for routing...")
	return &chiServer{
		port:             port,
		enablePlayground: enablePlayground,
		graphqlPath:      graphqlPath,
		graphqlHandler:   graphqlHandler,
		services:         services,
		version:          version,
		stage:            stage,
	}
}

type Middleware = func(next http.Handler) http.Handler

func (s *chiServer) Start() error {
	router := chi.NewRouter()
	router.Use(s.corsMiddleware)
	router.Get("/status", s.statusHandler)

	if s.enablePlayground {
		router.Handle("/", playground.Handler("Anime Skip", s.graphqlPath))
	}
	router.Route(s.graphqlPath, func(r chi.Router) {
		r.Use(s.ipMiddleware)
		r.Use(s.directivesMiddleware)
		r.Use(s.apiClientMiddleware)
		r.Handle("/", s.graphqlHandler.Handler)
	})

	log.I("Started server @ :%d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), router)
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

func (s *chiServer) apiClientMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientID := utils.Ptr(r.Header.Get("X-Client-ID"))
		if clientID == nil || strings.TrimSpace(*clientID) == "" {
			writeJson(rw, map[string]any{
				"errors": []map[string]any{{"message": "The X-Client-ID header must be passed"}},
			}, http.StatusOK)
			return
		}
		_, err := s.services.APIClientService.Get(r.Context(), internal.APIClientsFilter{
			ID: utils.Ptr(strings.TrimSpace(*clientID)),
		})

		if err != nil {
			writeJson(rw, map[string]any{
				"errors": []map[string]any{{"message": "Invalid X-Client-ID header, API client not found"}},
			}, http.StatusOK)
		} else {
			next.ServeHTTP(rw, r)
		}
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
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE")
		r.Header.Set("Access-Control-Allow-Origin", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Client-ID")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func writeJson(rw http.ResponseWriter, data any, status int) {
	rw.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	rw.Write(body)
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
