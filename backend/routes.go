package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/api/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/api/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/api/todos/{todoId}",
		TodoShow,
	},
}
