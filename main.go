package main

import (
	"clean-architecture/config"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//Load env file (.env)
	config.LoadEnv()
	//Connect db
	config.GetDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run("localhost:" + os.Getenv("PORT"))
}
