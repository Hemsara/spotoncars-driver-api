package routes

import (
	dvrController "spotoncars_server/controllers/drivers"
	"spotoncars_server/middleware"

	"github.com/gin-gonic/gin"
)

func DriverRoutes(router *gin.Engine) {
	protectedRoutesDriverApp := router.Group("/driverApp")
	protectedRoutesDriverApp.Use(middleware.DriverAuthenticationGuard)
	{
		driverAppRoutes := protectedRoutesDriverApp.Group("/drivers")
		driverAppRoutes.POST("/log", dvrController.LogDriver)
	}

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationGuard)
	{
		driverRoutes := protectedRoutes.Group("/drivers")
		{
			driverRoutes.GET("", dvrController.GetAllDrivers)
		}
	}
}
