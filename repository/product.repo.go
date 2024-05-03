package repository

import (
	"clean-architecture/model"
	"context"
	"gorm.io/gorm"
)

type ProductRepo interface {
	//GetCategories(cat model.Product) ([]model.Product, error)
	Create(ctx context.Context, model model.ProductCreate) error
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

func (repo *productRepo) Create(ctx context.Context, product model.ProductCreate) error {

	if err := repo.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}
