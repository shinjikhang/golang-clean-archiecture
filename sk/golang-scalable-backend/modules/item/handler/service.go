package handler

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

type ItemUseCase interface {
	GetItems(ctx context.Context, filter *model.Filter, pagination *common.Pagination) ([]model.TodoItem, error)
	GetItemById(ctx context.Context, id int) (*model.TodoItem, error)
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
	UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error
	DeleteItemById(ctx context.Context, id int) error
}

type itemService struct {
	useCase ItemUseCase
}
