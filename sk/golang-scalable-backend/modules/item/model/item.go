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

	ErrItemNotFound = errors.New("item not found")
	ErrTitleEmpty   = errors.New("Title cannot empty")

	ErrStatusInvalid = errors.New("Status invalid")
)

type TodoItem struct {
	common.SqlBase
	Title       string `json:"title" gorm:"column:title;'"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}
type TodoItemCreate struct {
	Id          int    `json:"id" gorm:"column:id" binding:"required"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}
type TodoItemUpdate struct {
	Title       string  `json:"title" gorm:"column:title"` //pointer cho phep update chuoi rong
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
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
