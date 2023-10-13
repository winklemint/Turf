package main

import (
	"net/http"
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

	r.Use(forbidHTMLExtension)

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)
<<<<<<< Updated upstream

	// Define a route that responds with "Hello, World!" when accessed
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
=======
	route.RegisterAdminPanelRoutes(r)
>>>>>>> Stashed changes

	r.Run(":8080")
}

func forbidHTMLExtension(c *gin.Context) {
	// Check if the URL path ends with ".html".
	if len(c.Request.URL.Path) > 5 && c.Request.URL.Path[len(c.Request.URL.Path)-5:] == ".html" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
		c.Abort()
		return
	}

	c.Next()
}
