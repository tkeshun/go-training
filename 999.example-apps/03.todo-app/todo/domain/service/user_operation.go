package service

import (
	"context"
	"todo/domain/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.User) error
	DeleteUserByID(ctx context.Context, userID model.UserID) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, userID model.UserID) (model.User, error)
}
