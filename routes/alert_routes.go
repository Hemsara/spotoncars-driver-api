package routes

import (
	alertController "spotoncars_server/controllers/notifications"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func AlertRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationGuard)
	{
		alertRoutes := protectedRoutes.Group("/alerts")
		{
			alertRoutes.POST("", alertController.SendNotification)
		}
	}
}
