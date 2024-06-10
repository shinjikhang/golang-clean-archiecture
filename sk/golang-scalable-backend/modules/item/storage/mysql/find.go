package mysql

import (
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	err := s.db.Where(cond).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
