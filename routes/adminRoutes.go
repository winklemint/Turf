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
		adminRoutes.PUT("update/admin", controllers.Package)
		// Slot
		adminRoutes.POST("add/slot", controllers.AddSlot)
		adminRoutes.POST("/update/slot/:id", controllers.UpdateSlot)
		adminRoutes.GET("/get/slot", controllers.GetAllSlot)
		//package
		adminRoutes.POST("add/package", controllers.AddPackage)
		adminRoutes.POST("/update/package/:id", controllers.UpdatePackage)
		adminRoutes.GET("/get/package", controllers.GetAllPackage)
		//booking
		adminRoutes.GET("/get/conform/booking", controllers.GetConfirmBooking)
		adminRoutes.POST("/update/confirm/booking/:id", controllers.UpdatecomfirmDetails)

	}
}
