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

	//r.Use(forbidHTMLExtension)

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

	r.Run(":8080")
}

// func forbidHTMLExtension(c *gin.Context) {
// 	// Check if the URL path ends with ".html".
// 	if len(c.Request.URL.Path) > 5 && c.Request.URL.Path[len(c.Request.URL.Path)-5:] == ".html" {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
// 		c.Abort()
// 		return
// 	}

// 	c.Next()
// }
