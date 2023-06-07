package middlewares

import (
	"GoGinStarter/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MaintenanceMode(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.Query("hash")
		if hash == "23a7711a-8133-4876-b7eb-dcd9e87a1613" {
			c.Next()
			return
		}
		if config.App.Maintenance == true {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "Service is temporarily unavailable due to maintenance",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
