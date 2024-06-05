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
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`           // neu co pointer se lay gia tri null insert vao db
	UpdatedAt   *time.Time `json:"updated_at,omitempty"` // neu co pointer se lay gia tri null insert vao db
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local"
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
			items.GET("", CreateItem(db)) // phai call () vi function nay return ve 1 function khac ( closure )
			items.GET("/:id")
			items.POST("")
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

	}
}
