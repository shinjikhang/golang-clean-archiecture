package dto

type Register struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	Mobile   string  `json:"mobile" binding:"required,min=10"`
	Bio      *string `json:"bio"`
}
