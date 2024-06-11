package gin_fw

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"clean-architecture/sk/golang-scalable-backend/modules/item/business"
	"clean-architecture/sk/golang-scalable-backend/modules/item/model"
	"clean-architecture/sk/golang-scalable-backend/modules/item/storage/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var queryString struct {
			//Search     string `form:"search"`
			common.Pagination
			model.Filter
		}

		//log request query
		fmt.Println(ctx.Request.URL.Query())

		if err := ctx.ShouldBindQuery(&queryString); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Pagination.Process()

		store := mysql.NewSQLStore(db)
		biz := business.NewGetListItemBusiness(store)

		items, err := biz.GetItems(ctx.Request.Context(), &queryString.Filter, &queryString.Pagination)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(items, queryString.Pagination, queryString.Filter))
	}
}
