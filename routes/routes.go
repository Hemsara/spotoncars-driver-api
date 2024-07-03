package routes

// routes/routes.go

import (
	controllers "spotoncars_server/controllers/bookings"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Apply middleware globally
	router.Use(middleware.AuthenticationGuard)

	// Define routes
	router.GET("/bookings/active", controllers.GetActiveBookings)
}
