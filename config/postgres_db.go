package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var dbDefault *gorm.DB

// sử dụng singleton pattern để tạo một connection duy nhất đến database
// khi ứng dụng lớn hơn thì không nên sử dụng singleton pattern
// thay vào đó nên sử dụng connection pool
func GetDB() *gorm.DB {
	if dbDefault == nil {
		initDb := initDB()
		if initDb != nil {
			dbDefault = initDb
		}
	}
	return dbDefault
}

func initDB() *gorm.DB {
	// Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	// Kết nối đến cơ sở dữ liệu PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db
}
