package infra

import (
	"todo/domain/model"
	"todo/domain/repo"
)

type UserRepo struct {
}

var _ repo.UserRepo = (*UserRepo)(nil)

func (u *UserRepo) Create(user model.User) {

}
