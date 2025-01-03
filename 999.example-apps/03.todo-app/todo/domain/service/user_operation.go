package service

import "todo/domain/model"

type UserService interface {
	RegisterUser(user model.User) error
	DeleteUserByID(userID model.UserID) error
	GetAllUsers() ([]model.User, error)
	GetUserByID() (model.User, error)
}
