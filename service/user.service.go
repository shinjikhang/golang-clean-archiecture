package service

import (
	"clean-architecture/model"
	"clean-architecture/repository"
	"context"
)

type User struct {
	repo repository.IUserRepo
}

func NewUser(repo repository.IUserRepo) *User {
	return &User{
		repo: repo,
	}
}

type IUserService interface {
	GetUsers(ctx context.Context) (users []model.User, err error)
}

func (s *User) GetUsers(ctx context.Context) (users []model.User, err error) {
	// Get user from repo
	users, err = s.repo.GetUsers(ctx)
	if err != nil {
		return users, err
	}
	return users, nil
}
