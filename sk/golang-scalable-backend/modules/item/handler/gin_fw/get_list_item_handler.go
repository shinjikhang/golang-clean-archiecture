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

func GetListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var queryString struct {
			Search     string `form:"search"`
			pagination common.Pagination
			filter     model.Filter
		}

		if err := ctx.ShouldBindQuery(&queryString); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.pagination.Process()

		store := mysql.NewSQLStore(db)
		biz := business.NewGetListItemBusiness(store)

		items, err := biz.GetItems(ctx.Request.Context(), &queryString.filter, &queryString.pagination)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(items, queryString.pagination, queryString.filter))
	}
}
