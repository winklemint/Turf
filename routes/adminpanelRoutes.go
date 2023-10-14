package routes

import (
	"net/http"
	"turf/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPanelRoutes(router *gin.Engine) {

	formGroup := router.Group("/panel")
	formGroup.Use(IsAuthenticated())

	formGroup.StaticFS("/", http.Dir("templates")) // Serve all files first

}



func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		//uri := c.Param("path")
		// Check the request path
		requestedPath := c.Request.URL.Path

		// If the requested path is "/panel/index.html," allow access without login
		if requestedPath == "/panel/" {
			//c.Next()
			return
		}

		// Check if the token is present in the request's cookies
		cookie, err := c.Cookie("Authorization")
		if err != nil || cookie == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Request": c,
			})
			c.Abort()
			return
		}

		middleware.RequireAdminAuth(c)

		// If the token is valid, continue
		c.Next()
	}
}
