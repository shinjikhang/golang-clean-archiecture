package gin_fw

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/business"
	"clean-architecture/sk/golang-scalable-backend/modules/item/storage/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Panic and recover test
		// demo panic and recover
		//go func() {
		//	defer common.Recovery()
		//
		//	var a []int
		//	log.Println(a[0])
		//}()

		var a []int
		log.Println(a[0])

		//var itemData model.TodoItem
		id, err := strconv.Atoi(ctx.Param("id"))

		log.Println("id", id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := mysql.NewSQLStore(db)
		biz := business.NewGetItemBusiness(store)

		data, err := biz.GetItemById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
