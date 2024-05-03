package service

import (
	dto "clean-architecture/dto/product"
	"clean-architecture/model"
	"clean-architecture/repository"

	"context"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *dto.CreateProductDto) error
}

type productService struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) *productService {
	return &productService{productRepo: productRepo}
}

func (service *productService) CreateProduct(ctx context.Context, productDto *dto.CreateProductDto) error {
	productData := model.ProductCreate{
		Title:       productDto.Title,
		Description: productDto.Description,
		Brand:       productDto.Brand,
		Price:       productDto.Price,
		SalePrice:   productDto.SalePrice,
		CategoryId:  productDto.CategoryId,
		Quantity:    productDto.Quantity,
		Images:      productDto.Images,
	}

	if err := service.productRepo.Create(ctx, productData); err != nil {
		return err
	}
	return nil
}
