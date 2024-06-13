package main

import (
	"clean-architecture/sk/golang-scalable-backend/middleware"
	"clean-architecture/sk/golang-scalable-backend/modules/item/handler/gin_fw"
	"clean-architecture/sk/golang-scalable-backend/modules/upload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL
	})
	//db = db.Debug() // Debug SQL
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)

	r := gin.Default()
	r.Use(middleware.Recover())

	//static
	r.Static("/uploads", "./uploads") // dùng để serve file tĩnh từ thư mục uploads (ví dụ: http://localhost:8080/uploads/abc.jpg)

	v1 := r.Group("/api/v1")
	{
		// upload route
		v1.POST("/upload", upload.Upload(db))
		items := v1.Group("/items")
		{
			items.GET("", gin_fw.GetListItem(db))
			items.GET("/:id", gin_fw.GetItem(db))
			items.POST("", gin_fw.CreateItem(db))
			items.PUT("/:id", gin_fw.UpdateItem(db))
			items.DELETE("/:id", gin_fw.DeleteItem(db))
		}
	}
	r.Run()
}
