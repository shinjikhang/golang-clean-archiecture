package business

import (
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

type GetItemStorage interface {
	//GetItem(ctx context.Context, id int) (*model.TodoItem, error)
	// thay vi truyen id, ta truyen vao 1 map[string]interface{} de co the truyen nhieu dieu kien khac nhau nhu id, name, status, ..., linh hoat hon
	// vi du: GetItem(ctx, map[string]interface{}{"id": 1})
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}
type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}
