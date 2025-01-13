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
	if err := u.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) DeleteUserByID(userID model.UserID) error {
	if err := u.userRepo.DeleteByID(userID); err != nil {
		return err
	}

	return nil
}
func (u *UserUseCase) GetAllUsers() ([]model.User, error) {
	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (u *UserUseCase) GetUserByID(userID model.UserID) (model.User, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
