package main

import (
	"fmt"
	"os"
	"turf/config"

	"turf/routes"
	route "turf/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	file, err := os.OpenFile("logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logrus.log")
		panic(err)
	}

	logrus.SetOutput(file)
	logrus.SetLevel(logrus.TraceLevel)

	r.Use(routes.ForbidHTMLExtension)

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)

	route.RegisterAdminPanelRoutes(r)
	//dashb0ard
	route.RegisterAdminPanelDashboard(r)
	//route.RegisterAdminPanelDashAdmin(r)

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
	route.RegisterAdminPanelUpdateBranchImage(r)
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

	// Details
	route.RegisterAdminPanelDetails(r)
	route.RegisterAdminPanelUpdateDetails(r)

	//Profile
	route.RegisterAdminPanelProfile(r)
	r.Run(":8080")
}
