package main

import (
	"turf/config"

	"turf/routes"
	route "turf/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.SyncDB()

}

func main() {
	//go controllers.Slot_go_rountine()
	// go controllers.MainCalendar()
	r := gin.Default()
	//r.GET("/proxy-react", frontend.ProxyHandlerReact)

	r.Use(routes.ForbidHTMLExtension)

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)

	route.RegisterAdminPanelRoutes(r)
	route.RegisterAdminPanelDashboard(r)

	//Carousel
	route.RegisterAdminPanelUpdatecarousel(r)
	route.RegisterAdminPanelAddCarousel(r)
	//Testimonial
	route.RegisterAdminPanelAddTestimonials(r)
	route.RegisterAdminPanelAllTestiMonials(r)
	route.RegisterAdminPanelUpdateTestiMonials(r)
	//Content
	route.RegisterAdminPanelUpdateContent(r)
	route.RegisterAdminPaneladdContent(r)
	route.RegisterAdminPanelOtherContent(r)
	route.RegisterAdminPanelUpdateOtherContent(r)
	// Branchs
	route.RegisterAdminPanelUpdatebranchs(r)
	route.RegisterAdminPanelAllBranch(r)
	route.RegisterAdminPanelCreateBranch(r)
	//Packages
	route.RegisterAdminPanelUpdatepackage(r)
	route.RegisterAdminPanelAddPackages(r)
	route.RegisterAdminPanelAllPackages(r)
	//Users
	route.RegisterAdminPanelAllUser(r)
	route.RegisterAdminPanelUpdateUser(r)
	route.RegisterAdminPanelAddUser(r)
	//Slots
	route.RegisterAdminPanelAllSlots(r)
	route.RegisterAdminPanelCreateSlots(r)
	//PSR
	route.RegisterAdminPanelPSR(r)
	//bkings
	route.RegisterAdminPanelAll_bookings(r)
	route.RegisterAdminPanelConfirmed_bookings(r)
	route.RegisterAdminPanelUpdatebookings(r)
	route.RegisterAdminPanelMultiBooking(r)

	//Remaining Amount
	route.RemainingAmountForAdminPanel(r)

	//Navbar
	route.RegisterAdminPanelNavbar(r)
	route.RegisterAdminPanelUpdateNavbar(r)

	//Social Icon
	route.RegisterAdminPanelSocialIcon(r)
	route.RegisterAdminPanelUpdateIcon(r)

	//staff
	route.RegisterAdminPanelAddstaff(r)
	route.RegisterAdminPanelAllstaff(r)
	route.RegisterAdminPanelUpdateStaff(r)

	r.Run(":8080")
}

// package main

// import (
// 	"turf/config"
// 	"turf/routes"
// 	route "turf/routes"
// 	"github.com/gin-gonic/gin"
// )

// func init() {
// 	config.LoadEnvVariables()
// 	config.ConnectToDB()
// 	config.SyncDB()
// }

// func main() {
// 	r := gin.Default()

// 	// CORS middleware for all domains
// 	r.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "600")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	})

// 	routes.RegisterAdminRoutes(r)
// 	route.RegisterUserRoutes(r)
// 	route.RegisterAdminPanelRoutes(r)
// 	route.RegisterAdminPanelDashboard(r)
// 	route.RegisterAdminPanelUpdatecarousel(r)
// 	route.RegisterAdminPanelAddCarousel(r)
// 	route.RegisterAdminPanelAddTestimonials(r)
// 	route.RegisterAdminPanelAllTestiMonials(r)
// 	route.RegisterAdminPanelUpdateTestiMonials(r)
// 	route.RegisterAdminPanelUpdateContent(r)
// 	route.RegisterAdminPaneladdContent(r)
// 	route.RegisterAdminPanelOtherContent(r)
// 	route.RegisterAdminPanelUpdateOtherContent(r)
// 	route.RegisterAdminPanelUpdatebranchs(r)
// 	route.RegisterAdminPanelAllBranch(r)
// 	route.RegisterAdminPanelCreateBranch(r)
// 	route.RegisterAdminPanelUpdatepackage(r)
// 	route.RegisterAdminPanelAddPackages(r)
// 	route.RegisterAdminPanelAllPackages(r)
// 	route.RegisterAdminPanelAllUser(r)
// 	route.RegisterAdminPanelUpdateUser(r)
// 	route.RegisterAdminPanelAddUser(r)
// 	route.RegisterAdminPanelAllSlots(r)
// 	route.RegisterAdminPanelCreateSlots(r)
// 	route.RegisterAdminPanelPSR(r)
// 	route.RegisterAdminPanelAll_bookings(r)
// 	route.RegisterAdminPanelConfirmed_bookings(r)
// 	route.RegisterAdminPanelUpdatebookings(r)
// 	route.RegisterAdminPanelMultiBooking(r)
// 	route.RemainingAmountForAdminPanel(r)
// 	route.RegisterAdminPanelNavbar(r)
// 	route.RegisterAdminPanelUpdateNavbar(r)
// 	route.RegisterAdminPanelSocialIcon(r)
// 	route.RegisterAdminPanelUpdateIcon(r)
// 	route.RegisterAdminPanelAddstaff(r)

// 	r.Run(":8080")
// }
