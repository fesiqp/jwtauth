package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fesiqp/jwtauth/api/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(h *handlers.Handler) *mux.Router {
	var routes = Routes{
		Route{
			"Index route",
			"GET",
			"/",
			h.Index,
		},
		Route{
			"Register User",
			"POST",
			"/register",
			h.RegisterUser,
		},
		Route{
			"Login User",
			"POST",
			"/login",
			h.Login,
		},
		Route{
			"Find User by Email",
			"GET",
			"/users/email/{email}",
			h.FindUserByEmail,
		},
		Route{
			"Find all Users",
			"GET",
			"/users",
			h.FindAllUsers,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Log(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Log(h http.Handler, name string) http.Handler {
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
