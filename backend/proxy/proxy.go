package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

}
