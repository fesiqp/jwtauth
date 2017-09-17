package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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

func JWTSignKey() ([]byte, error) {
	jwtString := os.Getenv("JWT_SIGN_KEY")
	if len(jwtString) == 0 {
		return []byte{}, fmt.Errorf("Login failed")
	}
	return []byte(jwtString), nil
}

func NewToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "jwtauth",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSignKey, err := JWTSignKey()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(jwtSignKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
	if auth.Scheme != "Bearer" && auth.Scheme != "Token" {
		return nil, fmt.Errorf("Only Bearer or Token allowed")
	}

	return jwt.Parse(auth.Token, keyFunc)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return JWTSignKey()
}
