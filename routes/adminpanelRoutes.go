package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPanelRoutes(router *gin.Engine) {

	formGroup := router.Group("/panel")

	formGroup.StaticFS("/", http.Dir("templates")) // Serve all files first

}
