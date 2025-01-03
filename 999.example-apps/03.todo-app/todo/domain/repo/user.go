package repo

import "todo/domain/model"

type UserRepo interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
	GetByID(userID model.UserID) (model.User, error)
	DeleteByID(userID model.UserID) error
}
