package routes

import (
	"turf/controllers"
	"turf/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {

	adminRoutes := router.Group("/admin")
	{
		//Admin and Staff
		adminRoutes.POST("/signup", controllers.AdminSignup)
		adminRoutes.POST("/login", controllers.AdminLogin)
		adminRoutes.PUT("/update", controllers.UpdateAdmin)
		adminRoutes.GET("/staff/get", controllers.AllStaff)
		adminRoutes.PATCH("/staff/update/:id", controllers.AdminUpdateById)
		adminRoutes.GET("/staff/get/:id", controllers.AdminGetById)
		adminRoutes.DELETE("/staff/delete/:id", controllers.AdminDelete)
		adminRoutes.GET("/profile", controllers.AdminProfile)
		adminRoutes.PATCH("update/profile", controllers.UpdateProfile)

		adminRoutes.GET("/data/login", controllers.GetLoggedAdmin)
		// Slot
		adminRoutes.POST("/add/slot", controllers.AddSlot)
		adminRoutes.POST("/update/slot/:id", controllers.UpdateSlot)
		adminRoutes.GET("/get/slot", controllers.GetAllSlot)
		adminRoutes.POST("/get/slot/:id", controllers.Get_Package)
		adminRoutes.DELETE("/delete/slot/:id", controllers.DeleteSlot)
		adminRoutes.POST("/get/avl/multi/slot", controllers.Get_Available_slots_Multi_Dates)

		//package
		adminRoutes.POST("/add/package", controllers.AddPackage)
		adminRoutes.PATCH("/update/package/:id", controllers.UpdatePackage)
		adminRoutes.GET("/get/package", controllers.GetAllPackage)
		adminRoutes.GET("/get/package/:id", controllers.GetAllPackageById)
		adminRoutes.DELETE("/delete/package/:id", controllers.DeletePackage)
		//booking
		adminRoutes.GET("/get/confirm/booking", controllers.Cnfrm_slots)
		adminRoutes.GET("/get/confirm/booking/top5", controllers.GetConfirmBookingTop5)
		// adminRoutes.GET("/get/confirm/booking/top5/:id", controllers.GetConfirmBookingTop5Super)
		adminRoutes.PATCH("/update/confirm/booking/:id", controllers.UpdatecomfirmDetails)
		adminRoutes.GET("/total/today/booking", controllers.Today_Total_Booking)
		adminRoutes.POST("/add/screenshot/:id", controllers.AdminAddScreenshot)
		adminRoutes.POST("/add/slot/:id", controllers.AddSlotForUser)
		adminRoutes.POST("/get/booking/date", controllers.GetAllDetailbydate)
		adminRoutes.POST("/remaining/payement/booking", controllers.RemainingPaymentForUser)
		adminRoutes.GET("/pending/bookings", controllers.Pending_bookings)
		adminRoutes.GET("/pending/bookings/:id", controllers.Pending_bookings_by_ID)
		adminRoutes.GET("/payments/:id", controllers.MultipleImages)

		adminRoutes.POST("/multi/bookings/:id", controllers.Multiple_slot_booking)

		//user Details
		adminRoutes.POST("/add/user", controllers.AddUser)
		adminRoutes.GET("/get/all/user", controllers.GetAllUsers)
		adminRoutes.PATCH("/update/user/:id", controllers.UpdateUserDetails)
		adminRoutes.GET("/get/user/:id", controllers.GetAllUsersById)
		adminRoutes.DELETE("/delete/user/:id", controllers.DeleteUser)
		adminRoutes.GET("user/count", controllers.CountUser)
		adminRoutes.GET("/count/user/monthly", controllers.GetMonthlyUsers)
		// adminRoutes.GET("/get/branch/name", controllers.Select_branch)

		//adminRoutes.GET("/pending/booking", controllers.Pending_bookings)
		//adminRoutes.GET("/partial/payments", controllers.Partial_payment)
		//adminRoutes.POST("/update/user/:id", controllers.UpdateUserDetails)
		adminRoutes.GET("/get/live/data", middleware.RequireAdminAuth, controllers.LiveUser)
		//Branch
		adminRoutes.POST("/add/branch", controllers.Add_Branch)
		//adminRoutes.PATCH("/update/branch/:id", controllers.Update_Branch)
		adminRoutes.GET("/get/branch", controllers.GET_All_Branch)
		adminRoutes.GET("/active/branch", controllers.ActiveBranch)
		adminRoutes.POST("set/id/branch", controllers.Get_IdBy_Branch_NAme)
		adminRoutes.GET("/get/branch/:id", controllers.GET_All_Branch_Id)
		adminRoutes.DELETE("/delete/branch/:id", controllers.Delete_Branch)
		adminRoutes.GET("/get/branch/image/:id", controllers.ImagesById)
		adminRoutes.PATCH("/update/branch/image/:id", controllers.UpdateImageById)
		adminRoutes.DELETE("/delete/branch/image/:id", controllers.DeleteImageById)
		adminRoutes.POST("/add/branch/image", controllers.AddImageForBranch)
		adminRoutes.GET("/get/image/:id", controllers.GetImageById)

		// adminRoutes.GET("/branch/image/active/:image", controllers.GetImageByImageName)

		adminRoutes.POST("/get/slot/by/day", controllers.Get_Slot_by_day)
		//tetsimonials
		adminRoutes.POST("/add/testimonials", controllers.Testimonials)
		adminRoutes.PATCH("/update/testimonials/:id", controllers.Upadte_Testimonials)
		adminRoutes.PATCH("/update/image/testimonials/:id", controllers.UpdateImageForTestimonials)
		adminRoutes.PATCH("/update/image/last/testimonials", controllers.UpdateImageForTestimonials2)
		adminRoutes.GET("/get/testimonials", controllers.AllTestimonials)
		adminRoutes.GET("/get/testimonial/:id", controllers.GETTestimonialsById)
		adminRoutes.GET("/get/testimonial/image/:id", controllers.GETTestimonialsimagesById)
		adminRoutes.DELETE("/delete/testimonial/:id", controllers.DeleteTestimonials)
		adminRoutes.GET("/testimonial/get/:branchid", controllers.GetTestimonialsBybranchId)

		//Content
		adminRoutes.POST("/content/add", controllers.AddContent)
		adminRoutes.GET("/content/get", controllers.GETContent)
		adminRoutes.PATCH("/content/update/:id", controllers.UpdateContent)
		adminRoutes.GET("/content/get/:id", controllers.GetContentById)
		adminRoutes.DELETE("/content/delete/:id", controllers.DeleteContent)
		adminRoutes.GET("/content/active", controllers.ActiveContent)

		//Carousel
		adminRoutes.POST("/carousel/add", controllers.AddImageForCarousel)
		adminRoutes.GET("/carousel/get", controllers.GetAllImageCarousel)
		adminRoutes.GET("/carousel/active", controllers.GetActiveImageCarousel)
		adminRoutes.PATCH("/carousel/upadte/:id", controllers.Upadtecarousel)
		adminRoutes.PATCH("/carousel/image/upadte/:id", controllers.UpadtecarouselImage)
		adminRoutes.DELETE("/delete/carousel/:id", controllers.DeleteCarousel)
		adminRoutes.GET("/get/image/active", controllers.GETCarouselActiveImages)
		adminRoutes.GET("/get/image/active/:id", controllers.GetCarouselimagesById)
		adminRoutes.GET("/get/slot/relationship", controllers.PSR_slots)

		//admin logout
		adminRoutes.POST("/logout", controllers.AdminLogout)

		//Navbar
		adminRoutes.POST("/navbar/add", controllers.AddNavbar)
		adminRoutes.GET("/navbar/get", controllers.GetAllNavbar)
		adminRoutes.GET("/navbar/active", controllers.GetActiveNavbar)
		adminRoutes.PATCH("/navbar/update/:id", controllers.UpadateNavbar)
		adminRoutes.GET("/navbar/get/:id", controllers.GetNavbarById)
		adminRoutes.DELETE("/navbar/delete/:id", controllers.DeleteNavbar)

		//Other content
		adminRoutes.POST("/heading/add", controllers.AddHeading)
		adminRoutes.GET("/heading/get", controllers.GetAllHeading)
		adminRoutes.GET("/heading/active", controllers.GetActiveHeading)
		adminRoutes.PATCH("/heading/update/:id", controllers.UpadateHeading)
		adminRoutes.GET("/heading/get/:id", controllers.GetHeadingById)
		adminRoutes.DELETE("/heading/delete/:id", controllers.DeleteHeading)

		//dashbard
		adminRoutes.GET("/total/revenue", controllers.Total_Revenue)
		adminRoutes.GET("/total/remaining/amount", controllers.Total_Remaining_amount)
		adminRoutes.GET("/total/sales", controllers.Total_Sales)
		adminRoutes.GET("/total/monthly/revenue", controllers.Total_Monthly_revenue)
		adminRoutes.GET("/sales/rati0", controllers.Graph_API)
		adminRoutes.GET("/package/list", controllers.PackageNameList)
		adminRoutes.GET("/total/slot", controllers.TotalSlot)
		//Social Icon
		adminRoutes.POST("/icon/add", controllers.AddIcon)
		adminRoutes.GET("/icon/get", controllers.GetAllIcon)
		adminRoutes.GET("/icon/active", controllers.GetActiveIcon)
		adminRoutes.PATCH("/icon/update/:id", controllers.UpadateIcon)
		adminRoutes.GET("/icon/get/:id", controllers.GetIconById)
		adminRoutes.DELETE("/icon/delete/:id", controllers.DeleteIcon)

		//TermsAndConditions
		adminRoutes.POST("/condition/add", controllers.AddTermsAndConditions)
		adminRoutes.GET("/condition/get", controllers.GetAllTermAndCondition)
		adminRoutes.GET("/condition/active", controllers.GetActiveTermAndCondition)
		adminRoutes.PATCH("/condition/update/:id", controllers.UpadateTermAndCondition)
		adminRoutes.GET("/condition/get/:id", controllers.GetTermAndConditionById)
		adminRoutes.DELETE("/condition/delete/:id", controllers.DeleteTermAndCondition)

	}

}
