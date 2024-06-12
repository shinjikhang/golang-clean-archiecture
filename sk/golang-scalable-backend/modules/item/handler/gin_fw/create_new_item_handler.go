package gin_fw

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/business"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"clean-architecture/sk/golang-scalable-backend/modules/item/storage/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Tầng handler này có nhiệm vụ nhận request, parse request (body), check header, gọi business logic và trả về response

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemData model.TodoItemCreate
		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := mysql.NewSQLStore(db)
		biz := business.NewCreateItemBusiness(store)

		// Lấy dữ liệu từ request và gọi business logic để tạo mới item trong database
		// (lưu ý: business logic không nên truy cập request hoặc response)
		if err := biz.CreateItem(ctx.Request.Context(), &itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))
	}
}
