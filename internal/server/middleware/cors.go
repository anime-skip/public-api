package middleware

import (
	"net/http"
)

// Prod config - TODO: Figure out origins for prod
var allowOrigin = "*"
var allowMethods = "GET, POST, OPTIONS"
var allowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

// Dev config
// func init() {
// 	if env.IS_DEV {
// 		allowOrigin = "*"
// 	}
// }

// Cors sets up the allowed controls for CORS and requests from browsers
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}
