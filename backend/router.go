package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupStatic(router *mux.Router) {
	var poopHandler http.HandlerFunc
	poopHandler = StaticHandler
	router.PathPrefix("/public/").Handler(Logger(poopHandler, "/public/"))
	router.PathPrefix("/static/").Handler(Logger(poopHandler, "/static/"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	setupStatic(router)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
