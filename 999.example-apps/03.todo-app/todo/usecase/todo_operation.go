package usecase

import (
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/service"
	"todo/domain/value"
)

type TodoUseCase struct {
	repoTodo repo.TodoRepo
}

var _ service.TodoService = (*TodoUseCase)(nil) // 左辺にinterface, 右辺に定義した型を置くことで、iterfaceを実装してるか確認できる。コンパイル時の型検査を活かすためにも、この処理は必須で書くべき

func (t *TodoUseCase) TodoGetAll() ([]model.Todo, error) {
	users, err := t.repoTodo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (t *TodoUseCase) TodoGetByUser(userID model.UserID) ([]model.Todo, error) {
	todos, err := t.repoTodo.GetByUser(userID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoUseCase) TodosCreate(todos []model.Todo) error {
	if err := t.repoTodo.Create(todos); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoComplete(todoID model.TodoID) error {
	if err := t.repoTodo.ChangeStatus(todoID, value.COMPLETED); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoReopen(todoID model.TodoID) error {
	if err := t.repoTodo.ChangeStatus(todoID, value.UNFINISHED); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoPriorityChange(todoID model.TodoID, priority value.TodoPriority) error {
	t.repoTodo.ChangePriority(todoID, priority)
	return nil
}

func (t *TodoUseCase) TagAddToTodo(todoID model.TodoID, tag model.Tag) error {
	if err := t.repoTodo.ChangeTag(todoID, tag); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TagRemoveFromTodo(todoID model.TodoID, tag model.Tag) error {
	if err := t.repoTodo.ChangeTag(todoID, tag); err != nil {
		return err
	}

	return nil
}

func (t *TodoUseCase) TodoFilter(priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error) {
	todos, err := t.repoTodo.GetByPriorityAndStatus(priority, status)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
