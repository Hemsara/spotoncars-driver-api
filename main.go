package main

import (
	"spotoncars_server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.InitDB()

}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
