package mysql

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
	"fmt"
)

func (s *sqlStore) GetItems(ctx context.Context, filter *model.Filter, pagination *common.Pagination, moreKey ...string) ([]model.TodoItem, error) {
	var items []model.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if status := f.Status; status != "" {
			fmt.Println("aa")
			db = db.Where("status = ?", status)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Table(model.TodoItem{}.TableName()).
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
