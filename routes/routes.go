package routes

import (
	controllers "spotoncars_server/controllers/bookings"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.Use(middleware.AuthenticationGuard)

	router.GET("/bookings/active", controllers.GetActiveBookings)
}
