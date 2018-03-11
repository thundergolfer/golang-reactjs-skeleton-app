package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
