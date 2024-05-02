package repository

import (
	"clean-architecture/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	//Register(user model.UserCreate) error
	//FindByEmail(email string) (*gorm.DB, model.User)
	GetCategories(cat model.Product) ([]model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db: db}
}

func (repo *productRepo) GetAll(cat model.Product) ([]model.Product, error) {
	category := []model.Product{}
	if err := repo.db.Order("category_id desc").Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
