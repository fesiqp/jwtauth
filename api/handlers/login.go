package handlers

import (
	"encoding/json"
	"net/http"

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
		http.Error(w, "Email or Password incorrect", http.StatusUnauthorized)
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
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
