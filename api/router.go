package main

import (
	"net/http"

	"github.com/fesiqp/jwtauth/api/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Secured     bool
}

type Routes []Route

func NewRouter(h *handlers.Handler) *mux.Router {
	var routes = Routes{
		Route{
			"Index route",
			"GET",
			"/",
			h.Index,
			false,
		},
		Route{
			"Register User",
			"POST",
			"/register",
			h.RegisterUser,
			false,
		},
		Route{
			"Login User",
			"POST",
			"/login",
			h.Login,
			false,
		},
		Route{
			"Find User by Email",
			"GET",
			"/users/email/{email}",
			h.FindUserByEmail,
			true,
		},
		Route{
			"Find User by Username",
			"GET",
			"/users/{username}",
			h.FindUserByUsername,
			true,
		},
		Route{
			"Find all Users",
			"GET",
			"/users",
			h.FindAllUsers,
			true,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Log(handler, route.Name)
		r := router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name)
		if route.Secured {
			r.Handler(handlers.ValidateTokenMiddleware(handler))
			continue
		}
		r.Handler(handler)
	}

	return router
}
