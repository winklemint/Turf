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
	r := gin.Default()

	routes.RegisterAdminRoutes(r)
	route.RegisterUserRoutes(r)
	route.RegisterAdminPanelRoutes(r)
	

	// Define a route that responds with the "login.html" template
	//r.LoadHTMLGlob("./templates/*.html")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// })

	r.Run(":8080")
}
