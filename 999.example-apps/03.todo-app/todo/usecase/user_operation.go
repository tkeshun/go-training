package usecase

import (
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/service"
)

type UserUseCase struct {
	userRepo repo.UserRepo
}

var _ service.UserService = (*UserUseCase)(nil)

func (u *UserUseCase) RegisterUser(user model.User) error {
	return nil
}

func (u *UserUseCase) DeleteUserByID(userID model.UserID) error {
	return nil
}
func (u *UserUseCase) GetAllUsers() ([]model.User, error) {
	return nil, nil
}
func (u *UserUseCase) GetUserByID() (model.User, error) {
	return model.User{}, nil
}
