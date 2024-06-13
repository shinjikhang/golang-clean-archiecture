package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"column:-"` // AmazonS3, GoogleCloud, Cloudinary
	Extension string `json:"extension,omitempty" gorm:"column:-"`
}

func (Image) TableName() string {
	return "images"
}
func (i *Image) FullFill(domain string) {
	i.Url = fmt.Sprintf("%s/%s", domain, i.Url)
}

// Dùng để chuyển đổi dữ liệu từ DB sang dạng JSON
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal data from database", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}

// Dùng để chuyển đổi dữ liệu từ JSON sang dạng DB
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}
