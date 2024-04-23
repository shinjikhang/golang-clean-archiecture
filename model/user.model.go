package model

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
	UserId   int     `json:"-" gorm:"primaryKey:auto_increment;column:user_id"` // ko cho phep parse json
	UserName string  `json:"-" gorm:"unique;column:user_name"`
	Name     string  `json:"name" gorm:"column:name;" db:"name"`
	Password string  `json:"password" gorm:"column:password;"`
	Email    string  `json:"email" gorm:"column:email;"`
	Bio      *string `json:"bio" gorm:"column:bio;"`
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
