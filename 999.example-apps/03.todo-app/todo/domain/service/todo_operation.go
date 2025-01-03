package service

import (
	"todo/domain/model"
	"todo/domain/value"
)

type TodoService interface {
	TodosCreate(todos []model.Todo) error
	TodoGetAll() ([]model.Todo, error)                       // viewにかえるかも
	TodoGetByUser(userID model.UserID) ([]model.Todo, error) // viewに変えるかも
	TodoComplete(todoID model.TodoID) error
	TodoReopen(todoID model.TodoID) error
	TodoPriorityChange(todoID model.TodoID, priority value.TodoPriority) error
	TagAddToTodo(todoID model.TodoID, tag model.Tag) error
	TagRemoveFromTodo(todoID model.TodoID, tag model.Tag) error
	TodoFilter(priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error)
}
