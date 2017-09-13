package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fesiqp/jwtauth/api/models"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	u := &body

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	m, err := models.NewUser(u.Username, u.Email, u.Password)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.DB.CreateUser(m)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	res := &User{
		Username: m.Username,
		Email:    m.Email,
	}

	json.NewEncoder(w).Encode(res)
}
