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

type LoginUser struct {
	Email    string
	Password string
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) FindUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	u, err := h.DB.FindUserByUsername(username)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := &User{
		Username: u.Username,
		Email:    u.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	u, err := h.DB.FindAllUsers()
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	users := make([]*User, len(u))
	for index, user := range u {
		users[index] = &User{Username: user.Username, Email: user.Email}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(users)

}
