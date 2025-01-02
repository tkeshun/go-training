package model

import "todo/domain/value"

type User struct {
	ID           UserID
	Name         UserName
	Email        UserEmail
	PasswordHash PasswordHash
	Role         value.Role
}

type UserID int64

func (i *UserID) validate() error {
	return nil
}

type UserName string

func (u *UserName) validate() error {
	return nil
}

type UserEmail string

func (e *UserEmail) validate() error {
	return nil
}

type PasswordHash string

func (p *PasswordHash) validate() error {
	return nil
}
