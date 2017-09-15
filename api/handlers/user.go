package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (h *Handler) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	u, err := h.DB.FindUserByEmail(email)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := &User{
		Username: u.Username,
		Email:    u.Email,
	}

	json.NewEncoder(w).Encode(res)
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	u, _ := h.DB.FindAllUsers()

	json.NewEncoder(w).Encode(u)

}
