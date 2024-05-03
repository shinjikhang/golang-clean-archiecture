package dto

type CreateProductDto struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Brand       string  `json:"brand" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	SalePrice   float32 `json:"sale_price"`
	CategoryId  uint    `json:"category_id" binding:"required"`
	Quantity    int     `json:"quantity"`
	Images      string  `json:"images" binding:"required"`
}

//type Register struct {
//	Name     string  `json:"name" binding:"required"`
//	Email    string  `json:"email" binding:"required,email"`
//	Password string  `json:"password" binding:"required,min=6"`
//	Mobile   string  `json:"mobile" binding:"required,min=10"`
//	Bio      *string `json:"bio"`
//}
