package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	pgconn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("pgx", pgconn)
	if err != nil {
		log.Fatalf("Openのエラー: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("接続確認の結果エラー: %v", err)
	}

	fetchUsers(context.Background(), db)
}

func fetchUsers(ctx context.Context, db *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		log.Fatalf("クエリエラー: %v", err)
	}
	defer rows.Close()

	fmt.Println("ユーザ一覧 (Context使用):")
	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatalf("行のスキャンエラー: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
}
