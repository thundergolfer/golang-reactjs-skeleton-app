package datastores

import (
	"github.com/thundergolfer/golang-reactjs-skeleton-app/backend/types"
)

type Datastore interface {
	ListTodos() types.Todos
	FindTodo(id string) types.Todo
	CreateTodo(t types.Todo) types.Todo
	DestroyTodo(id string) error
}
