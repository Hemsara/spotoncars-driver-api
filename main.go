package main

import (
	authController "spotoncars_server/controllers/auth"
	bookingController "spotoncars_server/controllers/bookings"
	dvrController "spotoncars_server/controllers/drivers"

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

	router.POST("/login", authController.LoginAdmin)

	router.Use(middleware.AuthenticationGuard)

	router.GET("/bookings/active", bookingController.GetActiveBookings)
	router.GET("/bookings/history", bookingController.GetBookingsHistory)

	router.GET("/drivers", dvrController.GetAllDrivers)

	router.Run()

}
