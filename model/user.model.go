package model

import "time"

const (
	USER_ROLE  = "user"
	ADMIN_ROLE = "admin"
)

//type RegisterUser struct {
//	FirstName string    `json:"firstName" validate:"required"`
//	LastName  string    `json:"lastName" validate:"required"`
//	UserName  string    `json:"userName" validate:"required"`
//	Email     string    `json:"email" validate:"required,email"`
//	Password  string    `json:"password" validate:"required,min=4"`
//	CreatedAt time.Time `validate:"required"`
//}
//
//func NewRegisterUser(firstName string, lastName string, userName string, email string, password string) *RegisterUser {
//	return &RegisterUser{FirstName: firstName, LastName: lastName, UserName: userName, Email: email, Password: password, CreatedAt: time.Now()}
//}

// Tạo struct để lưu thông tin của một bản ghi trong bảng "User"
type User struct {
	UserId   int     `json:"user_id" gorm:"primaryKey:auto_increment;type:uint64;column:user_id"` // ko cho phep parse json
	UserName string  `json:"user_name" gorm:"unique;column:user_name"`
	Name     string  `json:"name" gorm:"column:name;" db:"name"`
	Password string  `json:"password" gorm:"column:password;"`
	Email    string  `json:"email" gorm:"uniqueIndex;type:varchar(255);column:email;"`
	Bio      *string `json:"bio" gorm:"column:bio;"`
}

type UserCreate struct {
	UserId                int       `json:"-" gorm:"primaryKey:auto_increment;column:user_id"` // ko cho phep parse json
	UserName              string    `json:"-" gorm:"unique;column:user_name"`
	ActivationCode        string    `json:"-" gorm:"column:activation_code"`
	ActivationCodeExpired time.Time `json:"-" gorm:"column:activation_code_expired"`
	Name                  string    `json:"name" gorm:"column:name;" db:"name"`
	Password              string    `json:"password" gorm:"column:password;"`
	Email                 string    `json:"email" gorm:"column:email;"`
	Mobile                string    `json:"mobile" gorm:"column:mobile;"`
	Bio                   *string   `json:"bio" gorm:"column:bio;"`
	Role                  string    `json:"-" gorm:"column:role;"`
}

func (User) TableName() string {
	return "User"
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

// func (UserUpdate) TableName() string {
// 	return User{}.TableName()
// }
