package handlers

import "net/http"
import "encoding/json"
import "golang.org/x/crypto/bcrypt"

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
	res := &User{
		Username: user.Username,
		Email:    user.Email,
		Token:    "placeholder",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Authorization", "Bearer "+res.Token)
	json.NewEncoder(w).Encode(res)
}
