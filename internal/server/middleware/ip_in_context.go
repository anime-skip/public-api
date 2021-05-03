package middleware

import (
	"context"
	"net/http"

	"anime-skip.com/backend/internal/utils/constants"
)

// IPInContext puts the remote address in the request context
func IPInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constants.CTX_IP_ADDRESS, r.RemoteAddr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
