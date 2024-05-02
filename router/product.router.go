package router

import (
	logger "clean-architecture/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouter(db *gorm.DB, catRouter *gin.RouterGroup, logger *logger.Logger) {
	//var (
	//	catRepository repository.ProductCategoryRepo       = repository.NewProductCategoryRepo(db)
	//	catService    service.ProductCategoryService       = service.NewProductCategoryService(catRepository)
	//	catController controller.ProductCategoryController = controller.NewProductCategoryController(catService, logger)
	//)
	//catRouter.GET("", catController.GetAll)
}
