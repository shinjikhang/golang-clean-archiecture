package serializer

import "clean-architecture/model"

type CategoryResponse struct {
	CategoryId uint   `json:"category_id"`
	Title      string `json:"title"`
}

type CategorySerializer struct {
	ProductCategory model.ProductCategory
}

func (serializer *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		CategoryId: uint(serializer.ProductCategory.CategoryId),
		Title:      serializer.ProductCategory.Title,
	}
}

type CategoriesSerializer struct {
	Categories []model.ProductCategory
}

func (serializer *CategoriesSerializer) Response() []CategoryResponse {
	response := []CategoryResponse{}
	for _, category := range serializer.Categories {
		categorySerializer := CategorySerializer{ProductCategory: category}
		response = append(response, categorySerializer.Response())
	}
	return response
}
