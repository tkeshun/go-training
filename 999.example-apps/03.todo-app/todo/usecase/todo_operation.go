package usecase

import (
	"context"
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/service"
	"todo/domain/value"
)

type TodoUseCase struct {
	repoTodo repo.TodoRepo
	tx       Transaction
}

var _ service.TodoService = (*TodoUseCase)(nil) // 左辺にinterface, 右辺に定義した型を置くことで、iterfaceを実装してるか確認できる。コンパイル時の型検査を活かすためにも、この処理は必須で書くべき

func NewTodoService(repo repo.TodoRepo, tx Transaction) *TodoUseCase {
	return &TodoUseCase{
		repoTodo: repo,
		tx:       tx,
	}
}

func (t *TodoUseCase) TodoGetAll(ctx context.Context) ([]model.Todo, error) {
	var getUsers []model.Todo
	err := t.tx.RunInTransaction(ctx, func(ctx context.Context) error {
		users, err := t.repoTodo.GetAll(ctx)
		if err != nil {
			return err
		}
		getUsers = users
		return nil
	})
	if err != nil {
		return nil, err
	}

	return getUsers, nil
}

func (t *TodoUseCase) TodoGetByUser(ctx context.Context, userID model.UserID) ([]model.Todo, error) {
	todos, err := t.repoTodo.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoUseCase) TodosCreate(ctx context.Context, todos []model.Todo) error {
	if err := t.repoTodo.Create(ctx, todos); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoComplete(ctx context.Context, todoID model.TodoID) error {
	if err := t.repoTodo.ChangeStatus(ctx, todoID, value.COMPLETED); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoReopen(ctx context.Context, todoID model.TodoID) error {
	if err := t.repoTodo.ChangeStatus(ctx, todoID, value.UNFINISHED); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoPriorityChange(ctx context.Context, todoID model.TodoID, priority value.TodoPriority) error {
	t.repoTodo.ChangePriority(ctx, todoID, priority)
	return nil
}

func (t *TodoUseCase) TagAddToTodo(ctx context.Context, todoID model.TodoID, tag model.Tag) error {
	if err := t.repoTodo.ChangeTag(ctx, todoID, tag); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TagRemoveFromTodo(ctx context.Context, todoID model.TodoID, tag model.Tag) error {
	if err := t.repoTodo.ChangeTag(ctx, todoID, tag); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoFilter(ctx context.Context, priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error) {
	todos, err := t.repoTodo.GetByPriorityAndStatus(ctx, priority, status)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
