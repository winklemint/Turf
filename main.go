package main

import (
	"turf/config"
<<<<<<< Updated upstream
=======
	"turf/controllers"

>>>>>>> Stashed changes
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

	r.Run(":8080")
}
