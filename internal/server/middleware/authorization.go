package middleware

import (
	"context"
	"net/http"

	"anime-skip.com/backend/internal/utils/auth"
	"anime-skip.com/backend/internal/utils/constants"
)

// Authorization checks and loads authentication details
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("authorization")
		jwt, err := auth.ValidateAuthHeader(authHeader)
		if err != nil {
			ctx = context.WithValue(ctx, constants.CTX_AUTH_ERROR, err)
		}
		if jwt != nil {
			ctx = context.WithValue(ctx, constants.CTX_USER_ID, jwt["userId"])
			ctx = context.WithValue(ctx, constants.CTX_ROLE, jwt["role"])
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
