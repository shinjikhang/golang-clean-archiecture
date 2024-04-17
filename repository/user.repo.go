package repository

import (
	"clean-architecture/model"
	"context"
)

type IUserRepo interface {
	GetUsers(ctx context.Context) ([]model.User, error)
}

func (r *Repository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return []model.User{}, err
	}
	return users, nil
}
