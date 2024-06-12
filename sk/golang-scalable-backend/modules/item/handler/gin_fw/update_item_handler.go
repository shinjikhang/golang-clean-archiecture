package gin_fw

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/business"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"clean-architecture/sk/golang-scalable-backend/modules/item/storage/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		updateData := model.TodoItemUpdate{}

		if err := ctx.ShouldBind(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := mysql.NewSQLStore(db)
		biz := business.NewUpdateItemBusiness(store)

		if err := biz.UpdateItemById(ctx.Request.Context(), id, &updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
