package model

type ProductCategory struct {
	CategoryId int    `json:"category_id" gorm:"primaryKey:auto_increment;column:category_id"` // ko cho phep parse json
	Title      string `json:"title" gorm:"column:title"`
}

func (ProductCategory) TableName() string {
	return "ProductCategory"
}
