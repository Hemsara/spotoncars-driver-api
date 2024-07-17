package main

import (
	controllers "spotoncars_server/controllers/bookings"
	"spotoncars_server/initializers"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.InitDB()

}

func main() {
	router := gin.Default()
	router.Use(middleware.AuthenticationGuard)

	router.GET("/bookings/active", controllers.GetActiveBookings)
	router.GET("/bookings/history", controllers.GetBookingsHistory)

	router.Run()

}
