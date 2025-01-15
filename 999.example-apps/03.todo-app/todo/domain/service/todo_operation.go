package service

import (
	"context"
	"todo/domain/model"
	"todo/domain/value"
)

type TodoService interface {
	TodosCreate(ctx context.Context, todos []model.Todo) error
	TodoGetAll(ctx context.Context) ([]model.Todo, error)                         // viewにかえるかも
	TodoGetByUser(ctx context.Context, userID model.UserID) ([]model.Todo, error) // viewに変えるかも
	TodoComplete(ctx context.Context, todoID model.TodoID) error
	TodoReopen(ctx context.Context, todoID model.TodoID) error
	TodoPriorityChange(ctx context.Context, todoID model.TodoID, priority value.TodoPriority) error
	TagAddToTodo(ctx context.Context, todoID model.TodoID, tag model.Tag) error
	TagRemoveFromTodo(ctx context.Context, todoID model.TodoID, tag model.Tag) error
	TodoFilter(ctx context.Context, priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error)
}
