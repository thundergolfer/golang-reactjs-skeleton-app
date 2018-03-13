package datastores

import (
	"github.com/thundergolfer/12-factor/backend/types"
)

type Datastore interface {
	ListTodos() types.Todos
	FindTodo(id string) types.Todo
	CreateTodo(t types.Todo) types.Todo
	DestroyTodo(id string) error
}
