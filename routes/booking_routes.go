package routes

import (
	bookingController "spotoncars_server/controllers/bookings"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationGuard)
	{
		bookingRoutes := protectedRoutes.Group("/bookings")
		{
			bookingRoutes.GET("/active", bookingController.GetActiveBookings)
			bookingRoutes.GET("/history", bookingController.GetBookingsHistory)
		}
	}
}
