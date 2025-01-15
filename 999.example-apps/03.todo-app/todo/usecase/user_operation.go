package usecase

import (
	"context"
	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/service"
)

type UserUseCase struct {
	userRepo repo.UserRepo
}

var _ service.UserService = (*UserUseCase)(nil)

func NewUserService(useRepo repo.UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: useRepo,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user model.User) error {
	if err := u.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) DeleteUserByID(ctx context.Context, userID model.UserID) error {
	if err := u.userRepo.DeleteByID(ctx, userID); err != nil {
		return err
	}

	return nil
}
func (u *UserUseCase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (u *UserUseCase) GetUserByID(ctx context.Context, userID model.UserID) (model.User, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
