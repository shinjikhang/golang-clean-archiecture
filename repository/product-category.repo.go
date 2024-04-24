package repository

import (
	"clean-architecture/model"
	"gorm.io/gorm"
)

type ProductCategoryRepo interface {
	//Register(user model.UserCreate) error
	//FindByEmail(email string) (*gorm.DB, model.User)
	GetCategories(cat model.ProductCategory) ([]model.ProductCategory, error)
}

type productCategoryRepo struct {
	db *gorm.DB
}

func NewProductCategoryRepo(db *gorm.DB) *productCategoryRepo {
	return &productCategoryRepo{db: db}
}

func (repo *productCategoryRepo) GetCategories(cat model.ProductCategory) ([]model.ProductCategory, error) {
	category := []model.ProductCategory{}
	if err := repo.db.Order("category_id desc").Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
