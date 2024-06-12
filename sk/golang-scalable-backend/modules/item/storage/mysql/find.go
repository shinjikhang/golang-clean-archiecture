package mysql

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	err := s.db.Where(cond).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
