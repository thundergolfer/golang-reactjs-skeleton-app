package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	var poopHandler http.HandlerFunc
	poopHandler = StaticHandler
	router.PathPrefix("/public/").Handler(poopHandler)
	router.PathPrefix("/static/").Handler(poopHandler)

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
