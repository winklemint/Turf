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
		adminRoutes.PUT("/update", controllers.UpdateAdmin)
		// Slot
		adminRoutes.POST("add/slot", controllers.AddSlot)
		adminRoutes.POST("/update/slot/:id", controllers.UpdateSlot)
		adminRoutes.GET("/get/slot", controllers.GetAllSlot)
		//package
		adminRoutes.POST("add/package", controllers.AddPackage)
		adminRoutes.POST("/update/package/:id", controllers.UpdatePackage)
		adminRoutes.GET("/get/package", controllers.GetAllPackage)
		//booking
		adminRoutes.GET("/get/confirm/booking", controllers.GetConfirmBooking)
		adminRoutes.POST("/update/confirm/booking/:id", controllers.UpdatecomfirmDetails)
		adminRoutes.POST("/add/branch", controllers.Add_Branch)
		//user Details
		adminRoutes.GET("/get/all/user", controllers.GetAllUsers)
		adminRoutes.POST("/update/user/:id", controllers.UpdateUserDetails)

	}
}
