package main

import (
	"spotoncars_server/initializers"
	"spotoncars_server/middleware"
	"spotoncars_server/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.InitDB()

}

func main() {
	router := gin.New()

	router.Use(middleware.AuthenticationGuard)
	routes.SetupRoutes(router)

	router.Run()

}
