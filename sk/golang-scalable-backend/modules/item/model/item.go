package model

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"errors"
	"strings"
)

var (
	StatusPending = "pending"
	StatusDone    = "done"
	StatusDeleted = "deleted"

	ErrItemNotFound  = errors.New("item not found")
	ErrTitleEmpty    = errors.New("Title cannot empty")
	ErrStatusInvalid = errors.New("Status invalid")
	ErrItemIsDeleted = errors.New("Item is deleted")
)

const (
	EntityName = "TodoItem"
)

type TodoItem struct {
	common.SqlBase
	Title       string        `json:"title" gorm:"column:title;'"`
	Description string        `json:"description" gorm:"column:description;"`
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"` // dùng con trỏ để có thể null -> null khi không có ảnh
}
type TodoItemCreate struct {
	Id          int    `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}
type TodoItemUpdate struct {
	Title       string        `json:"title" gorm:"column:title"` //pointer cho phep update chuoi rong
	Description *string       `json:"description" gorm:"column:description"`
	Status      *string       `json:"status" gorm:"column:status"`
	Image       *common.Image `json:"image" gorm:"column:image;"` // dùng con trỏ để có thể null -> null khi không có ảnh
}

const tblName = "todo_items"

func (TodoItem) TableName() string       { return tblName }
func (TodoItemCreate) TableName() string { return tblName }
func (TodoItemUpdate) TableName() string { return tblName }

func (i *TodoItemCreate) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleEmpty
	}

	return nil
}
