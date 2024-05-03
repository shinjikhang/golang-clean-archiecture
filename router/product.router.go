package router

import (
	"clean-architecture/controller"
	logger "clean-architecture/pkg"
	"clean-architecture/repository"
	"clean-architecture/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouter(db *gorm.DB, productRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		productRepository repository.ProductRepo       = repository.NewProductRepo(db)
		productService    service.ProductService       = service.NewProductService(productRepository)
		productController controller.ProductController = controller.NewProductController(productService, logger)
	)
	productRouter.POST("", productController.CreateProduct)
}
