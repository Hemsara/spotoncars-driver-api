package main

import (
	"spotoncars_server/controllers"
	"spotoncars_server/initializers"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.InitDB()

}

func main() {
	router := gin.New()

	router.Use(middleware.AuthenticationGuard)

	router.GET("/bookings/active", controllers.GetActiveBookings)
	router.Run()

}
