package main

import (
	"clean-architecture/config"
	"clean-architecture/pkg"
	"clean-architecture/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//Load env file (.env)
	config.LoadEnv()
	//Connect db
	db := config.GetDB()
	// logger
	log := logger.NewLogger()

	routerGin := gin.Default()
	router.RootRoute(db, routerGin, log)
	routerGin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	routerGin.Run("localhost:" + os.Getenv("PORT"))
}
