package service

import (
	"clean-architecture/model"
	"clean-architecture/repository"
)

type ProductCategoryService interface {
	GetCategories(cat model.ProductCategory) ([]model.ProductCategory, error)
}

type productCategoryService struct {
	productCategoryRepo repository.ProductCategoryRepo
}

func NewProductCategoryService(productCategoryRepo repository.ProductCategoryRepo) *productCategoryService {
	return &productCategoryService{productCategoryRepo: productCategoryRepo}
}

func (service *productCategoryService) GetCategories(cat model.ProductCategory) ([]model.ProductCategory, error) {
	categories, err := service.productCategoryRepo.GetCategories(cat)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
