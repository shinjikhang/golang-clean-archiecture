package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id,primaryKey"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"created_at"`           // neu co pointer se lay gia tri null insert vao db
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"updated_at"` // neu co pointer se lay gia tri null insert vao db
}

type TodoItemCreate struct {
	Id          int    `json:"id" gorm:"column:id,primaryKey"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string       { return "todo_items" }
func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root@tcp(localhost:3306/social-todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)
	// GIN HTTP
	r := gin.Default()
	//Item router
	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", GetListItem(db)) // phai call () vi function nay return ve 1 function khac ( closure )
			items.GET("/:id")
			items.POST("", CreateItem(db))
			items.PUT("/:id")
			items.DELETE("/:id")
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		itemData := TodoItem{}
		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&itemData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": itemData.Id,
		})
	}
}

func GetListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		itemData := TodoItemCreate{}
		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&itemData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": itemData.Id,
		})
	}
}
