package routes

import (
	authController "spotoncars_server/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.LoginAdmin)
	}
}
