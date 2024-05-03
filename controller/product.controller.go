package controller

import (
	dto "clean-architecture/dto/product"
	logger "clean-architecture/pkg"
	"clean-architecture/service"
	"clean-architecture/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	CreateProduct(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	logger         *logger.Logger
}

func NewProductController(
	productService service.ProductService, logger *logger.Logger,
) *productController {
	return &productController{productService: productService, logger: logger}
}

func (controller *productController) CreateProduct(context *gin.Context) {
	var productDto dto.CreateProductDto
	if err := context.ShouldBindJSON(&productDto); err != nil {
		context.JSON(http.StatusInternalServerError, util.GetErrorResponse(err.Error()))
		return
	}

	if err := controller.productService.CreateProduct(context, &productDto); err != nil {
		context.JSON(http.StatusInternalServerError, util.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	//if err := controller.productService.CreateProduct(context, &product); err != nil {
	//	context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
	//	controller.logger.Error().Err(err).Msg("")
	//	return
	//}

	context.JSON(http.StatusOK, util.GetResponse([]int{}))
}
