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
		adminRoutes.POST("/add/slot", controllers.AddSlot)
		adminRoutes.POST("/update/slot/:id", controllers.UpdateSlot)
		adminRoutes.GET("/get/slot", controllers.GetAllSlot)

		//package
		adminRoutes.POST("/add/package", controllers.AddPackage)
		adminRoutes.POST("/update/package/:id", controllers.UpdatePackage)
		adminRoutes.GET("/get/package", controllers.GetAllPackage)
		//booking
		adminRoutes.GET("/get/confirm/booking", controllers.GetConfirmBooking)
		adminRoutes.GET("/get/confirm/booking/top5", controllers.GetConfirmBookingTop5)
		adminRoutes.POST("/update/confirm/booking/:id", controllers.UpdatecomfirmDetails)

		adminRoutes.POST("/add/screenshot/:id", controllers.AdminAddScreenshot)
		adminRoutes.POST("/add/slot/:id", controllers.AddSlotForUser)
		adminRoutes.GET("/get/booking/date", controllers.GetAllDetailbydate)

		//user Details
		//adminRoutes.GET("/get/all/user", controllers.GetAllUsers)
		//adminRoutes.POST("/get/user/:id", controllers.UpdateUserDetails)
		// adminRoutes.GET("/get/branch/name", controllers.Select_branch)

		//adminRoutes.GET("/pending/booking", controllers.Pending_bookings)
		//adminRoutes.GET("/partial/payments", controllers.Partial_payment)
		//adminRoutes.POST("/update/user/:id", controllers.UpdateUserDetails)
		//adminRoutes.POST("/get/live/data", controllers.LiveUser)
		//Branch
		adminRoutes.POST("/add/branch", controllers.Add_Branch)
		adminRoutes.POST("/update/branch/:id", controllers.Update_Branch)
		adminRoutes.GET("/get/branch", controllers.GET_All_Branch)

		adminRoutes.POST("/get/slot/by/day", controllers.Get_Slot_by_day)
		//tetsimonials
		adminRoutes.POST("/add/testimonials", controllers.Testimonials)
		adminRoutes.PUT("/update/testimonials/:id", controllers.Upadte_TestiMonilas)
		adminRoutes.GET("/get/testimonials", controllers.AllTestimonials)
		adminRoutes.POST("/get/slot/:id", controllers.Get_Package)

	}
}
