package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func routes(app *App) Routes {
	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			app.Index,
		},
		Route{
			"TodoIndex",
			"GET",
			"/api/todos",
			app.TodoIndex,
		},
		Route{
			"TodoCreate",
			"POST",
			"/api/todos",
			app.TodoCreate,
		},
		Route{
			"TodoShow",
			"GET",
			"/api/todos/{todoId}",
			app.TodoShow,
		},
		Route{
			"TodoDelete",
			"DELETE",
			"/api/todos/{todoId}",
			app.TodoDelete,
		},
	}

	return routes
}
