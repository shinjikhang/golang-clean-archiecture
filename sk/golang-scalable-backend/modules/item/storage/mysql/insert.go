package mysql

import (
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"context"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	//ctx dùng để truyền giá trị giữa các hàm mà không cần thông qua tham số truyền vào, và không cần biết hàm đó ở đâu
	//và dùng để debug xem dữ liệu truyền vào hàm là gì khi gặp lỗi hoặc cần kiểm tra dữ liệu truyền vào hàm
	return s.db.Create(&data).Error
}
