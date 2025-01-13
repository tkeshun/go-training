package model

import (
	"time"
	"todo/domain/value"
)

type Todo struct {
	ID          TodoID
	Title       TodoTitle
	Description string
	Status      value.TodoStatus
	Priority    value.TodoPriority
	tags        []Tag
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoID int64

type TodoTitle string
