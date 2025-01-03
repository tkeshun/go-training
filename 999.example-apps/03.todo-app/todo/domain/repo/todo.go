package repo

import (
	"todo/domain/model"
	"todo/domain/value"
)

type TodoRepo interface {
	GetAll() ([]model.Todo, error)
	Create(todos []model.Todo) error
	GetByUser(model.UserID) ([]model.Todo, error)
	ChangeStatus(todoID model.TodoID, status value.TodoStatus) error
	ChangePriority(todoID model.TodoID, priority value.TodoPriority) error
	ChangeTag(todoID model.TodoID, tag model.Tag) error
	GetByPriorityAndStatus(priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error)
}
