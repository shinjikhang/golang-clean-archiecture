package business

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

type ListItemStorage interface {
	GetItems(ctx context.Context,
		filter *model.Filter,
		pagination *common.Pagination,
		moreKey ...string,
	) ([]model.TodoItem, error)
}
type listItemBusiness struct {
	store ListItemStorage
}

func NewGetListItemBusiness(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (b *listItemBusiness) GetItems(ctx context.Context, filter *model.Filter, pagination *common.Pagination, moreKey ...string) ([]model.TodoItem, error) {
	data, err := b.store.GetItems(ctx, filter, pagination)

	if err != nil {
		return nil, err
	}

	return data, nil
}
