package main

import (
	"clean-architecture/config"
	"clean-architecture/middleware"
	logger "clean-architecture/pkg"
	"clean-architecture/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//Load env file (.env)
	config.LoadEnv()
	//Connect db
	db := config.GetDB()
	// logger
	log := logger.NewLogger()

	routerGin := gin.Default()

	// use middleware
	routerGin.Use(middleware.CORSMiddleware())

	router.RootRoute(db, routerGin, log)
	routerGin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	routerGin.Run("localhost:" + os.Getenv("PORT"))
}
