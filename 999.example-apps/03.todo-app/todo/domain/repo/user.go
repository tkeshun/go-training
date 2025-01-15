package repo

import (
	"context"
	"todo/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, userID model.UserID) (model.User, error)
	DeleteByID(ctx context.Context, userID model.UserID) error
}
