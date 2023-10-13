package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPanelRoutes(router *gin.Engine) {
	formGroup := router.Group("/panel")
	{
		// Serve files from the "form" directory with the "/form" path prefix
		formGroup.StaticFS("/", http.Dir("templates"))
	}
}
