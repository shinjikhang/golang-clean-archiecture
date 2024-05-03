package model

import "time"

//type CreateProduct struct {
//	ProductID   uuid.UUID `validate:"required"`
//	Name        string    `validate:"required,gte=0,lte=255"`
//	Description string    `validate:"required,gte=0,lte=5000"`
//	Price       float64   `validate:"required,gte=0"`
//	InventoryId int64     `validate:"required,gt=0"`
//	Count       int32     `validate:"required,gt=0"`
//	CreatedAt   time.Time `validate:"required"`
//}

type Product struct {
	ProductId   int        `json:"-" gorm:"primaryKey:AUTO_INCREMENT;column:product_id"`
	Title       string     `json:"title" gorm:"column:title"`
	Slug        string     `json:"-" gorm:"column:slug"`
	Description string     `json:"description" gorm:"column:description"`
	Brand       string     `json:"brand" gorm:"column:brand"`
	Price       float32    `json:"price" gorm:"column:price"`
	SalePrice   float32    `json:"sale_price" gorm:"column:sale_price"`
	CategoryId  uint       `json:"category_id" gorm:"column:category_id"`
	Quantity    int        `json:"quantity" gorm:"column:quantity"`
	Sold        uint       `json:"-" gorm:"column:sold"`
	Images      string     `json:"images" gorm:"column:images"`
	Color       string     `json:"color" gorm:"column:color"`
	Rating      uint       `json:"-" gorm:"column:rating"`
	CreatedAt   *time.Time // Automatically managed by GORM for creation time
	UpdatedAt   *time.Time // Automatically managed by GORM for update time
}

type ProductCreate struct {
	ProductId   int     `json:"-" gorm:"primaryKey:AUTO_INCREMENT;column:product_id"`
	Title       string  `json:"title" gorm:"column:title"`
	Slug        string  `json:"slug" gorm:"column:slug"`
	Description string  `json:"description" gorm:"column:description"`
	Brand       string  `json:"brand" gorm:"column:brand"`
	Price       float32 `json:"price" gorm:"column:price"`
	SalePrice   float32 `json:"sale_price" gorm:"column:sale_price"`
	CategoryId  uint    `json:"category_id" gorm:"column:category_id"`
	Quantity    int     `json:"quantity" gorm:"column:quantity"`
	Sold        uint    `json:"-" gorm:"column:sold"`
	Images      string  `json:"images" gorm:"column:images"`
	Color       string  `json:"-" gorm:"column:color"`
	Rating      uint    `json:"-" gorm:"column:rating"`
}

func (Product) TableName() string {
	return "Product"
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}
