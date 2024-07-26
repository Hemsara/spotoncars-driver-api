package main

import (
	"spotoncars_server/initializers"
	"spotoncars_server/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.InitDB()
	initializers.InitRedis()

}

func main() {
	router := gin.Default()

	routes.AuthRoutes(router)
	routes.DriverRoutes(router)
	routes.BookingRoutes(router)
	routes.AlertRoutes(router)

	router.Run()
}
