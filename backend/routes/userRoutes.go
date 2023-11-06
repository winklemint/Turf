package routes

import (
	"turf/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/signup", controllers.Signup)
		userRoutes.POST("/verify/otp", controllers.VerifyOTPhandler)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.POST("/booking", controllers.Booking)
		//userRoutes.POST("/available/slot", controllers.AvailableSlot)
		userRoutes.PUT("/update", controllers.UpdateUser)
		userRoutes.POST("/uplad", controllers.Screenshot)
		userRoutes.GET("/get/detail", controllers.GetAllDetail)
		userRoutes.GET("/get/booking/detail", controllers.GetBookingDetail)
		userRoutes.POST("/get/avl/slots", controllers.Get_Available_slots)

	}

}
