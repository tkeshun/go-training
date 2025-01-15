package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	schemaName     = "group_a"
	tableTodos     = schemaName + ".Todos"
	tableUserTodos = schemaName + ".UserTodos"
	tableTags      = schemaName + ".Tags"
	tableTodoTags  = schemaName + ".TodoTags"
	tableUsers     = schemaName + ".Users"
)

var pgconn *pgxpool.Pool

func GetConn() (*pgxpool.Pool, error) {
	var connErr error
	if pgconn != nil {
		_, connErr = NewDBConnection()
	}

	if connErr != nil {
		return nil, connErr
	}
	return pgconn, nil
}

// DATABASE_URL =　postgresql://postgres:postgres@127.0.0.1:5432/postgres
// DATABASE_URL=postgresql://postgres:postgres@127.0.0.1:5432/postgres go run main.go
func NewDBConnection() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create database connection pool: %w", err)
	}

	pgconn = pool
	return pool, nil
}

type PgxTransactionManager struct {
}

func NewPgxTransactionManager(pool *pgxpool.Pool) *PgxTransactionManager {
	return &PgxTransactionManager{}
}

type txContextKey struct{}

func (tm *PgxTransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// コネクションの取得
	pool, err := GetConn()
	if err != nil {
		return err
	}
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// トランザクション開始
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// ビジネスロジックを実行
	ctxWithTx := context.WithValue(ctx, txContextKey{}, tx)
	err = fn(ctxWithTx)
	if err != nil {
		// エラー時にロールバック
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return fmt.Errorf("rollback failed: %v, original error: %w", rollbackErr, err)
		}
		return err
	}

	// コミット
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func getTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(txContextKey{}).(pgx.Tx)

	if !ok {
		return nil, fmt.Errorf("transaction not found in context")
	}
	return tx, nil
}
