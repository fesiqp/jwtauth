package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fesiqp/jwtauth/api/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Log(h.Index, "Index")).Methods("GET")
	router.HandleFunc("/register", Log(h.RegisterUser, "Register User")).Methods("POST")
	router.HandleFunc("/user/email/{email}", Log(h.FindUserByEmail, "Find user by email")).Methods("GET")

	return router
}

func Log(h http.HandlerFunc, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "[ROUTE] ", log.LstdFlags)
		start := time.Now()

		h.ServeHTTP(w, r)

		logger.Printf(
			"%-5s%-50s\t%-30s\t%-10s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
