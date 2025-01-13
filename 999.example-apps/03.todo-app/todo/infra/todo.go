package infra

import (
	"context"
	"strings"
	"time"
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/value"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepo struct {
	conn *pgxpool.Pool
}

type todo struct {
	ID          int64
	Title       string
	Description string
	Status      string
	Priority    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var _ repo.TodoRepo = (*TodoRepo)(nil)

const (
	schemaName     = "group_a"
	tableTodos     = schemaName + ".Todos"
	tableUserTodos = schemaName + ".UserTodos"
	tableTags      = schemaName + ".Tags"
	tableTodoTags  = schemaName + ".TodoTags"
)

func NewTodoRepo(conn *pgxpool.Pool) repo.TodoRepo {
	return &TodoRepo{conn: conn}
}

func (t *TodoRepo) GetAll(ctx context.Context) ([]model.Todo, error) {
	var sb strings.Builder
	sb.WriteString("SELECT id, title, description, status, priority, created_at, updated_at FROM ")
	sb.WriteString(tableTodos)
	query := sb.String()

	rows, err := t.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, model.Todo{
			ID:          model.TodoID(todo.ID),
			Title:       model.TodoTitle(todo.Title),
			Description: todo.Description,
			Status: func() value.TodoStatus {
				status, err := value.TodoStatusString(todo.Status)
				if err != nil {
					return value.UNDEFINED
				}

				return status
			}(),
			Priority: func() value.TodoPriority {
				pr, err := value.TodoPriorityString(todo.Priority)
				if err != nil {
					return value.LOW
				}

				return pr
			}(),
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}

func (t *TodoRepo) Create(ctx context.Context, todos []model.Todo) error {
	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(tableTodos)
	sb.WriteString(" (title, description, status, priority) VALUES ($1, $2, $3, $4)")
	query := sb.String()

	for _, mTodo := range todos {
		todo := todo{
			Title:       string(mTodo.Title),
			Description: mTodo.Description,
			Status:      mTodo.Status.String(),
			Priority:    mTodo.Priority.String(),
		}
		_, err := t.conn.Exec(ctx, query, todo.Title, todo.Description, todo.Status, todo.Priority)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TodoRepo) GetByUser(ctx context.Context, userID model.UserID) ([]model.Todo, error) {
	var sb strings.Builder
	sb.WriteString("SELECT t.id, t.title, t.description, t.status, t.priority, t.created_at, t.updated_at ")
	sb.WriteString("FROM ")
	sb.WriteString(tableTodos)
	sb.WriteString(" t JOIN ")
	sb.WriteString(tableUserTodos)
	sb.WriteString(" ut ON t.id = ut.todo_id WHERE ut.user_id = $1")
	query := sb.String()

	rows, err := t.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}

func (t *TodoRepo) ChangeStatus(ctx context.Context, todoID model.TodoID, status value.TodoStatus) error {
	var sb strings.Builder
	sb.WriteString("UPDATE ")
	sb.WriteString(tableTodos)
	sb.WriteString(" SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2")
	query := sb.String()

	_, err := t.conn.Exec(ctx, query, status, todoID)
	if err != nil {
		return err
	}
	return nil
}

// ChangePriority updates the priority of a todo
func (t *TodoRepo) ChangePriority(ctx context.Context, todoID model.TodoID, priority value.TodoPriority) error {
	var sb strings.Builder
	sb.WriteString("UPDATE ")
	sb.WriteString(tableTodos)
	sb.WriteString(" SET priority = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2")
	query := sb.String()

	_, err := t.conn.Exec(ctx, query, priority, todoID)
	if err != nil {
		return err
	}
	return nil
}

// ChangeTag updates the tag of a todo
func (t *TodoRepo) ChangeTag(ctx context.Context, todoID model.TodoID, tag model.Tag) error {
	var sb strings.Builder

	// Check if the tag exists
	sb.WriteString("SELECT id FROM ")
	sb.WriteString(tableTags)
	sb.WriteString(" WHERE name = $1")
	tagQuery := sb.String()
	var tagID int64
	err := t.conn.QueryRow(ctx, tagQuery, tag.Name).Scan(&tagID)
	if err != nil {
		return err
	}

	// Update the TodoTags table
	sb.Reset()
	sb.WriteString("INSERT INTO ")
	sb.WriteString(tableTodoTags)
	sb.WriteString(" (todo_id, tag_id) VALUES ($1, $2) ON CONFLICT (todo_id, tag_id) DO NOTHING")
	query := sb.String()
	_, err = t.conn.Exec(ctx, query, todoID, tagID)
	if err != nil {
		return err
	}

	return nil
}

// GetByPriorityAndStatus retrieves todos filtered by priority and status
func (t *TodoRepo) GetByPriorityAndStatus(ctx context.Context, priority value.TodoPriority, status value.TodoStatus) ([]model.Todo, error) {
	var sb strings.Builder
	sb.WriteString("SELECT id, title, description, status, priority, created_at, updated_at FROM ")
	sb.WriteString(tableTodos)
	sb.WriteString(" WHERE priority = $1 AND status = $2")
	query := sb.String()

	rows, err := t.conn.Query(ctx, query, priority, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}
