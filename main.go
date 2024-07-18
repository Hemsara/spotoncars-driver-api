package main

import (
	authController "spotoncars_server/controllers/auth"
	bookingController "spotoncars_server/controllers/bookings"
	dvrController "spotoncars_server/controllers/drivers"
	alertController "spotoncars_server/controllers/notifications"

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

	// Authentication routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.LoginAdmin)
	}

	// Apply middleware to protected routes
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationGuard)
	{
		// Booking routes
		bookingRoutes := protectedRoutes.Group("/bookings")
		{
			bookingRoutes.GET("/active", bookingController.GetActiveBookings)
			bookingRoutes.GET("/history", bookingController.GetBookingsHistory)
		}

		// Driver routes
		driverRoutes := protectedRoutes.Group("/drivers")
		{
			driverRoutes.GET("", dvrController.GetAllDrivers)
			driverRoutes.POST("log", dvrController.LogDriver)

		}

		alertRoutes := protectedRoutes.Group("/alerts")
		{
			alertRoutes.POST("", alertController.SendNotification)
		}
	}

	router.Run()
}
