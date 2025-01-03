package service

import "todo/domain/model"

type userService interface {
	UserRegister(user model.User) error
	UserDelete(userID model.UserID) error
}
