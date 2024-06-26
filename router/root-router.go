package router

import (
	logger "clean-architecture/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine, logger *logger.Logger) {
	//router.Static("/media", "/media")
	apiRouter := router.Group("/api/v1")
	//postRouter := apiRouter.Group("/posts")
	//PostRoute(db, postRouter, logger)
	//commentRouter := apiRouter.Group("/posts/:postId/comments")
	//CommentRoute(db, commentRouter, logger)

	/// --------------------------------Category-------------------------------------
	categoryRouter := apiRouter.Group("/categories")
	ProductCategoryRouter(db, categoryRouter, logger)

	/// --------------------------------Product-------------------------------------
	procductRouter := apiRouter.Group("/product")
	ProductRouter(db, procductRouter, logger)

	authRouter := apiRouter.Group("/auth")
	AuthRoute(db, authRouter, logger)
}
