package common

import "time"

type SqlBase struct {
	ID        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
