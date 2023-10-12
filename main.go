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

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)

	// Define a route that responds with "Hello, World!" when accessed
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.Run(":8080")
}
