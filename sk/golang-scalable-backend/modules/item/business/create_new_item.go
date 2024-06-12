package business

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}
type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}

func (b *createItemBusiness) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrValidation(err)
	}

	if err := b.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
