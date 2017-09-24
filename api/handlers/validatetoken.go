package handlers

import (
	"log"
	"net/http"
	"os"
)

func ValidateTokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger := log.New(os.Stdout, "[ROUTE] ", log.LstdFlags)

		token, err := ValidateToken(r)
		if err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logger.Println(err)
			http.Error(w, "Token is invalid", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
