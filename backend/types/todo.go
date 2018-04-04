package types

import (
	"time"
)

type Todo struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
