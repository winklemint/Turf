package routes

import (
	"turf/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/signup", controllers.AdminSignup)
		adminRoutes.POST("/login", controllers.AdminLogin)
		adminRoutes.POST("add/slot", controllers.AddSlot)
		adminRoutes.POST("add/package", controllers.AddPackage)
		adminRoutes.PUT("update/package", controllers.Package)
	}
}
