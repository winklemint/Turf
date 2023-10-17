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

	//r.Use(forbidHTMLExtension)

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)
	route.RegisterAdminPanelRoutes(r)
	route.RegisterAdminPanelDashboard(r)
	route.RegisterAdminPanelCreateBranch(r)

	// r.LoadHTMLGlob("templates/*.tmpl")
	// // Define a route that uses the header and footer templates
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "main.tmpl", gin.H{
	// 		// "title": "My Page",
	// 	})
	// })

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
