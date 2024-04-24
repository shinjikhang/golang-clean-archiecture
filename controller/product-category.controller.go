package controller

import (
	"clean-architecture/model"
	logger "clean-architecture/pkg"
	"clean-architecture/serializer"
	"clean-architecture/service"
	"clean-architecture/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductCategoryController interface {
	GetAll(context *gin.Context)
}

type productCategoryController struct {
	categoryService service.ProductCategoryService
	logger          *logger.Logger
}

func NewProductCategoryController(
	categoryService service.ProductCategoryService, logger *logger.Logger,
) *productCategoryController {
	return &productCategoryController{categoryService: categoryService, logger: logger}
}

func (controller *productCategoryController) GetAll(context *gin.Context) {
	var catModel model.ProductCategory
	categories, err := controller.categoryService.GetCategories(catModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, util.GetErrorResponse(err.Error()))
		return
	}
	serializer := serializer.CategoriesSerializer{Categories: categories}
	context.JSON(http.StatusOK, util.GetResponse(serializer.Response()))
}
