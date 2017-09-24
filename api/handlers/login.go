package handlers

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

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
		http.Error(w, "Email or password incorrect", http.StatusUnauthorized)
		return
	}

	token, err := NewToken()
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := &User{
		Username: user.Username,
		Email:    user.Email,
		Token:    CheckActive(user.Email, token),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

var activeUsers = make(map[string]string) // [email]token

// CheckActive validates if an user is already authenticated
// If so, the JWT is re-sent, otherwise, a new JWT is generated.
func CheckActive(email string, token string) string {
	if tkn, ok := activeUsers[email]; ok {
		jwtToken, _ := jwt.Parse(tkn, keyFunc)
		if jwtToken.Valid {
			return tkn
		}
		activeUsers[email] = token
		return token
	}
	activeUsers[email] = token
	return token
}
