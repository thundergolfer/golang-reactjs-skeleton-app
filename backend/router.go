package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupStatic(router *mux.Router) {
	var handler http.HandlerFunc
	handler = StaticHandler
	router.PathPrefix("/public/").Handler(Logger(handler, "/public/"))
	router.PathPrefix("/static/").Handler(Logger(handler, "/static/"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	setupStatic(router)
	config := newConfig()
	app := newApp(config)

	routes := routes(app)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
