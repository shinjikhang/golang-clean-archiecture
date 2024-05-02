package serializer

import "clean-architecture/model"

type CategoryResponse struct {
	CategoryId uint   `json:"category_id"`
	Title      string `json:"title"`
}

type CategorySerializer struct {
	ProductCategory model.ProductCategory
}

type CategoriesSerializer struct {
	ProductCategories []model.ProductCategory
}

func (serializer *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		CategoryId: uint(serializer.ProductCategory.CategoryId),
		Title:      serializer.ProductCategory.Title,
	}
}

func (serializer *CategoriesSerializer) Response() []CategoryResponse {
	response := []CategoryResponse{}
	for _, category := range serializer.ProductCategories {
		categorySerializer := CategorySerializer{ProductCategory: category}
		response = append(response, categorySerializer.Response())
	}
	return response
}
