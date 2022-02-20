package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/log"
)

type chiServer struct {
	port             int
	enablePlayground bool
	graphqlPath      string
	graphqlHandler   internal.GraphQLHandler
	authenticator    internal.Authenticator
}

func NewChiServer(
	port int,
	enablePlayground bool,
	graphqlPath string,
	graphqlHandler internal.GraphQLHandler,
	authenticator internal.Authenticator,
) internal.Server {
	log.D("Using Chi for routing...")
	return &chiServer{
		port:             port,
		enablePlayground: enablePlayground,
		graphqlPath:      graphqlPath,
		graphqlHandler:   graphqlHandler,
		authenticator:    authenticator,
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
		r.Use(s.authorizerMiddleware)
		r.Handle("/", s.graphqlHandler.Handler)
	})

	log.I("Started server @ :%d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), router)
}

func (s *chiServer) statusHandler(rw http.ResponseWriter, r *http.Request) {
	status := internal.ApiStatus{
		Version:       config.VERSION,
		Status:        "RUNNING",
		Playground:    s.enablePlayground,
		Introspection: s.graphqlHandler.EnableIntrospection,
	}
	writeJson(rw, status, http.StatusOK)
}

func (s *chiServer) authorizerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithAuthenticator(ctx, s.authenticator)
		ctx = context.WithAuthToken(ctx, getAuthToken(r))
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

func writeJson(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	rw.Write(body)
	rw.WriteHeader(status)
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
