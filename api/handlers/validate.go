package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type Authorization struct {
	Scheme string
	Token  string
}

func NewAuthorization(authSlice []string) *Authorization {
	return &Authorization{
		Scheme: authSlice[0],
		Token:  authSlice[1],
	}
}

func ValidateTokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "[ROUTE] ", log.LstdFlags)

		token, err := ValidateToken(r)
		if err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		switch token.Valid {
		case true:
			h.ServeHTTP(w, r)
		default:
			fmt.Fprintln(w, "Token is invalid")
		}
	})
}

func ValidateToken(r *http.Request) (*jwt.Token, error) {
	authString := r.Header.Get("Authorization")
	if len(authString) == 0 {
		return nil, fmt.Errorf("Authorization Header is missing")
	}

	authHeaderSlice := strings.Split(authString, " ")
	if len(authHeaderSlice) != 2 {
		return nil, fmt.Errorf("Invalid Authorization Header")
	}

	auth := NewAuthorization(authHeaderSlice)
	if auth.Scheme != "Bearer" {
		return nil, fmt.Errorf("Only Bearer or Token allowed")
	}

	return jwt.Parse(auth.Token, keyFunc)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return JWTSignKey()
}
