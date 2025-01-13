package repo

import (
	"context"
	"todo/domain/model"
	"todo/domain/value"
)

type TodoRepo interface {
	GetAll(ctx context.Context) ([]model.Todo, error)
	Create(ctx context.Context, todos []model.Todo) error
	GetByUser(ctx context.Context, userID model.UserID) ([]model.Todo, error)
	ChangeStatus(ctx context.Context, todoID model.TodoID, status value.TodoStatus) error
	ChangePriority(ctx context.Context, todoID model.TodoID, priority value.TodoPriority) error
	ChangeTag(ctx context.Context, todoID model.TodoID, tag model.Tag) error
	GetByPriorityAndStatus(ctx context.Context, priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error)
}
