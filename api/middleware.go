package api

import (
	"net/http"
)

// jwt auth middleware

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, err := ValidateToken(tokenStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.Header.Set("username", claims.Username)
		next.ServeHTTP(w, r)
	})
}
