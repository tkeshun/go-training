package infra

import (
	"context"
	"fmt"
	"strings"
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/value"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	conn *pgxpool.Pool
}

var _ repo.UserRepo = (*UserRepo)(nil)

func NewUserRepo(conn *pgxpool.Pool) repo.UserRepo {
	return &UserRepo{conn: conn}
}

type UserDto struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Role         string
}

func (u *UserRepo) Create(ctx context.Context, user model.User) error {
	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(tableUsers)
	sb.WriteString(" (name, email, password_hash) VALUES ($1, $2, $3)")
	query := sb.String()

	insUser := UserDto{
		ID:           int64(user.ID),
		Name:         string(user.Name),
		Email:        string(user.Email),
		PasswordHash: string(user.PasswordHash),
		Role:         user.Role.String(),
	}

	_, err := u.conn.Exec(ctx, query, insUser.ID, insUser.Name, insUser.Email, insUser.PasswordHash, insUser.Role)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]model.User, error) {

	query := "SELECT id, name, email, password_hash, role FROM " + schemaName + tableUsers
	rows, err := u.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user UserDto

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role); err != nil {
			return nil, err
		}

		users = append(users, model.User{
			ID:           model.UserID(user.ID),
			Name:         model.UserName(user.Name),
			Email:        model.UserEmail(user.Email),
			PasswordHash: model.PasswordHash(user.PasswordHash),
			Role: func() value.Role {
				role, err := value.RoleString(user.Role)
				if err != nil {
					return value.REGULAR_USER
				}
				return role
			}(),
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (u *UserRepo) GetByID(ctx context.Context, userID model.UserID) (model.User, error) {
	query := "SELECT id, name, email, password_hash, role FROM Users WHERE id = $1"
	var user UserDto
	var role string
	if err := u.conn.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &role); err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:           model.UserID(user.ID),
		Name:         model.UserName(user.Name),
		Email:        model.UserEmail(user.Email),
		PasswordHash: model.PasswordHash(user.PasswordHash),
		Role: func() value.Role {
			role, err := value.RoleString(user.Role)
			if err != nil {
				return value.REGULAR_USER
			}
			return role
		}(),
	}, nil
}

func (u *UserRepo) DeleteByID(ctx context.Context, userID model.UserID) error {
	query := "DELETE FROM" + schemaName + tableUsers + " WHERE id = $1"
	cmd, err := u.conn.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("non delete user. Not Found user by ID")
	}
	return nil
}
