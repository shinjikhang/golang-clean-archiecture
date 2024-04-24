package router

import (
	"clean-architecture/controller"
	logger "clean-architecture/pkg"
	"clean-architecture/repository"
	"clean-architecture/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductCategoryRouter(db *gorm.DB, catRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		catRepository repository.ProductCategoryRepo       = repository.NewProductCategoryRepo(db)
		catService    service.ProductCategoryService       = service.NewProductCategoryService(catRepository)
		catController controller.ProductCategoryController = controller.NewProductCategoryController(catService, logger)
	)
	catRouter.GET("", catController.GetAll)
}
