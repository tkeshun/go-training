package infra

import (
	"context"
	"fmt"
	"os"
	"testing"
	"todo/domain/model"
	"todo/domain/value"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	// テスト用のDB接続を準備
	var err error
	testPool, err = connectDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	// テスト実行
	code := m.Run()

	// クリーンアップ
	testPool.Close()

	os.Exit(code)
}

func connectDB() (*pgxpool.Pool, error) {
	connAddr := "postgres://postgres:postgres@localhost:5432/postgres"
	pool, err := pgxpool.New(context.Background(), connAddr)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func Test_todoRepo_GetAll(t *testing.T) {
	// テストデータを挿入
	_, err := testPool.Exec(context.Background(), "INSERT INTO group_a.Todos (title, description, status, priority) VALUES ($1, $2, $3, $4)", "Test Todo", "Test Description", "UNFINISHED", "MEDIUM")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewTodoRepo(testPool)
	tests := []struct {
		name    string
		want    []model.Todo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			want: []model.Todo{
				{
					Title:       "Test Todo",
					Description: "Test Description",
					Status:      value.UNFINISHED,
					Priority:    value.MEDIUM,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetAll(context.Background())
			if !tt.wantErr(t, err, fmt.Sprintf("GetAll test %s", tt.name)) {
				return
			}

			for i := range got {
				got[i].ID = 0
				got[i].CreatedAt = tt.want[i].CreatedAt
				got[i].UpdatedAt = tt.want[i].UpdatedAt
			}

			assert.Equal(t, tt.want, got, "GetAll() returned unexpected result")
		})
	}
	testPool.Exec(context.Background(), "DELETE FROM group_a.Todos WHERE title = $1", "Test Todo")
}

func Test_todoRepo_Create(t *testing.T) {
	repo := NewTodoRepo(testPool)
	tests := []struct {
		name    string
		todos   []model.Todo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			todos: []model.Todo{
				{
					Title:       "New Todo",
					Description: "New Description",
					Status:      value.UNFINISHED,
					Priority:    value.HIGH,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Create(context.Background(), tt.todos)
			if !tt.wantErr(t, err, fmt.Sprintf("Create test %s", tt.name)) {
				return
			}
		})
	}
}

func Test_todoRepo_GetByUser(t *testing.T) {
	// テストデータを挿入
	_, err := testPool.Exec(context.Background(), `
		INSERT INTO group_a.Todos (title, description, status, priority) VALUES ('User Todo', 'User Description', 'UNFINISHED', 'LOW');
		INSERT INTO group_a.UserTodos (user_id, todo_id) VALUES (1, 1);
	`)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewTodoRepo(testPool)
	tests := []struct {
		name    string
		userID  model.UserID
		want    []model.Todo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "success",
			userID: 1,
			want: []model.Todo{
				{
					ID:          1,
					Title:       "User Todo",
					Description: "User Description",
					Status:      value.UNFINISHED,
					Priority:    value.LOW,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByUser(context.Background(), tt.userID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByUser test %s", tt.name)) {
				return
			}
			assert.Equal(t, tt.want, got, "GetByUser() returned unexpected result")
		})
	}
}

func Test_todoRepo_ChangeStatus(t *testing.T) {
	// テストデータを挿入
	_, err := testPool.Exec(context.Background(), "INSERT INTO group_a.Todos (id, title, description, status, priority) VALUES (1, 'Change Status Todo', 'Description', 'UNFINISHED', 'MEDIUM')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewTodoRepo(testPool)
	tests := []struct {
		name    string
		todoID  model.TodoID
		status  value.TodoStatus
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			todoID:  1,
			status:  value.COMPLETED,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.ChangeStatus(context.Background(), tt.todoID, tt.status)
			if !tt.wantErr(t, err, fmt.Sprintf("ChangeStatus test %s", tt.name)) {
				return
			}
		})
	}
}

func Test_todoRepo_ChangePriority(t *testing.T) {
	// テストデータを挿入
	_, err := testPool.Exec(context.Background(), "INSERT INTO group_a.Todos (id, title, description, status, priority) VALUES (1, 'Change Priority Todo', 'Description', 'UNFINISHED', 'MEDIUM')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewTodoRepo(testPool)
	tests := []struct {
		name     string
		todoID   model.TodoID
		priority value.TodoPriority
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "success",
			todoID:   1,
			priority: value.HIGH,
			wantErr:  assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.ChangePriority(context.Background(), tt.todoID, tt.priority)
			if !tt.wantErr(t, err, fmt.Sprintf("ChangePriority test %s", tt.name)) {
				return
			}
		})
	}
}

func Test_todoRepo_GetByPriorityAndStatus(t *testing.T) {
	// テストデータを挿入
	_, err := testPool.Exec(context.Background(), "INSERT INTO group_a.Todos (title, description, status, priority) VALUES ('Priority Status Todo', 'Description', 'UNFINISHED', 'MEDIUM')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewTodoRepo(testPool)
	tests := []struct {
		name     string
		priority value.TodoPriority
		status   value.TodoStatus
		want     []model.Todo
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "success",
			priority: value.MEDIUM,
			status:   value.UNFINISHED,
			want: []model.Todo{
				{
					ID:          1,
					Title:       "Priority Status Todo",
					Description: "Description",
					Status:      value.UNFINISHED,
					Priority:    value.MEDIUM,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByPriorityAndStatus(context.Background(), tt.priority, tt.status)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByPriorityAndStatus test %s", tt.name)) {
				return
			}
			assert.Equal(t, tt.want, got, "GetByPriorityAndStatus() returned unexpected result")
		})
	}
}
