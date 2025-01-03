package model

import (
	"todo/domain/value"
)

type Todo struct {
	ID          TodoID
	Title       TodoTitle
	Description string
	Status      value.TodoStatus
	Priority    value.TodoPriority
	tags        []Tag
}

type TodoID int64

type TodoTitle string
