package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"SystemInfo",
		"GET",
		"/",
		GetSystemInfo,
	},
	Route{
		"RandomUser",
		"GET",
		"/random",
		GetRandomUser,
	},
	Route{
		"UserInfo",
		"GET",
		"/user/{userId}",
		GetUserInfo,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
