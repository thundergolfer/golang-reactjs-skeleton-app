package datastores

import (
	"fmt"

	"github.com/satori/go.uuid"
	"github.com/thundergolfer/12-factor/backend/types"
)

type InMemoryStorer struct {
	todos     types.Todos
	currentId int
}

func NewInMemoryStorer() *InMemoryStorer {
	storer := InMemoryStorer{
		currentId: 1, // 0 is considered an invalid ID
	}

	// Create some seed data
	storer.CreateTodo(types.Todo{
		Text: "Test TODO 1",
	})
	storer.CreateTodo(types.Todo{
		Text: "Test TODO 2",
	})
	return &storer
}

func (s *InMemoryStorer) ListTodos() types.Todos {
	return s.todos
}

func (s *InMemoryStorer) FindTodo(id string) types.Todo {
	for _, t := range s.todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return types.Todo{}
}

func (s *InMemoryStorer) CreateTodo(t types.Todo) types.Todo {
	s.currentId += 1
	t.Id = uuid.NewV4().String()
	s.todos = append(s.todos, t)
	return t
}

func (s *InMemoryStorer) DestroyTodo(id string) error {
	for i, t := range s.todos {
		if t.Id == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %s to delete", id)
}
