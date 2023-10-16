package routes

import (
	"fmt"
	"net/http"
	"turf/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPanelRoutes(router *gin.Engine) {

	router.LoadHTMLGlob("templates/*.html")

	formGroup := router.Group("/panel")
	formGroup.Use(IsAuthenticated())

	formGroup.StaticFS("/", http.Dir("templates")) // Serve all files first

}

func RegisterAdminPanelDashboard(router *gin.Engine) {
	// Load HTML templates for the admin panel
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelCreateBranch(router *gin.Engine) {
	// Load HTML templates for the admin panel
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/create/branch", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createbranch.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

// Serve all files first

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		//uri := c.Param("path")
		// Check the request path
		requestedPath := c.Request.URL.Path

		// If the requested path is "/panel/index.html," allow access without login
		if requestedPath == "/panel/" {
			c.Next()
			return
		}

		// Check if the token is present in the request's cookies
		cookie, err := c.Cookie("Authorization")

		fmt.Println("Authorization Cookie Value:", cookie)

		if err != nil || cookie == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Request": c,
			})
			c.Abort()
			return
		}

		middleware.RequireAdminAuth(c)

		c.Next()
	}
}
