package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"created_at"`           // neu co pointer se lay gia tri null insert vao db
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"updated_at"` // neu co pointer se lay gia tri null insert vao db
}

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Pagination) Process() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 || p.Limit > 50 {
		p.Limit = 12
	}
}

type TodoItemCreate struct {
	Id          int    `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}
type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"` //pointer cho phep update chuoi rong
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

const tblName = "todo_items"

func (TodoItem) TableName() string       { return tblName }
func (TodoItemCreate) TableName() string { return tblName }
func (TodoItemUpdate) TableName() string { return tblName }

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL
	})

	//db = db.Debug() // Debug SQL

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
			items.GET("/:id", GetItem(db))
			items.POST("", CreateItem(db))
			items.PUT("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
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

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemData TodoItemCreate
		id, err := strconv.Atoi(ctx.Param("id"))

		log.Println("id", id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&itemData).Error; err != nil { // First lay 1 row, 1 hoac nhieu thi lay Find
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": itemData,
		})
	}
}

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		updateData := TodoItemUpdate{}

		if err := ctx.ShouldBind(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		deleteStatus := "Deleted"
		if err := db.Where("id = ?", id).Updates(&TodoItemUpdate{Status: &deleteStatus}).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func GetListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var pagination Pagination

		if err := ctx.ShouldBindQuery(&pagination); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		pagination.Process()

		var items []TodoItem

		if err := db.
			Model(&TodoItem{}).
			Count(&pagination.Total).
			Offset((pagination.Page - 1) * pagination.Limit).
			Limit(pagination.Limit).
			Order("id desc").
			Find(&items).
			Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":       items,
			"pagination": pagination,
		})
	}
}
