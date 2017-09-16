package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Email    string
	Password string
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginUser := &LoginUser{}

	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	user, err := h.DB.FindUserByEmail(loginUser.Email)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, "Email or Password incorrect", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, "Email or Password incorrect", http.StatusUnauthorized)
		return
	}

	token, err := NewToken(loginUser)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := &User{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func JWTSignKey() ([]byte, error) {
	jwtString := os.Getenv("JWT_SIGN_KEY")
	if len(jwtString) == 0 {
		return []byte{}, fmt.Errorf("Login failed")
	}
	return []byte(jwtString), nil
}

func NewToken(user *LoginUser) (string, error) {
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
