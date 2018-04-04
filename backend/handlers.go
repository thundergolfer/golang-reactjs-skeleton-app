package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/thundergolfer/golang-reactjs-skeleton-app/backend/datastores"
	"github.com/thundergolfer/golang-reactjs-skeleton-app/backend/types"
)

func newApp(c config) *App {
	a := App{}
	ctx := context.Background() // TODO: creating context and passing to storer is probably wrong

	switch c.datastoreType {
	case "local":
		a.datastore = datastores.NewInMemoryStorer()
	case "googlecloudstorage":
		a.datastore = datastores.NewGoogleCloudStorer(c.projectID, c.googleCloudStorageBucketName, ctx)
	default:
		panic(fmt.Sprintf("Unrecognised storer type: %s", c.datastoreType))
	}

	return &a
}

type App struct {
	datastore datastores.Datastore
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:] // strip leading slash
	if path == "" {
		path = "index.html"
	}
	if !strings.HasPrefix(path, "public") {
		path = "public/" + path
	}

	log.Println(AssetDir("public"))
	if bs, err := Asset(path); err != nil {
		log.Warn(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(w, reader)
	}
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	if bs, err := Asset("public/index.html"); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(w, reader)
	}
}

func (app *App) TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(app.datastore.ListTodos()); err != nil {
		panic(err)
	}
}

func (app *App) TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todo := app.datastore.FindTodo(vars["todoID"])
	if todo.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"text":"New Todo"}' http://localhost:8080/todos
*/
func (app *App) TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo types.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := app.datastore.CreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func (app *App) TodoDelete(w http.ResponseWriter, r *http.Request) {
	todoID := path.Base(r.URL.Path)

	// if err != nil {
	// 	log.Warn(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	err := app.datastore.DestroyTodo(todoID)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
