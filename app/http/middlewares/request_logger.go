package middlewares

import (
	"GoGinStarter/internal/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestLogger(log log.Log) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(fmt.Sprintf("[%s] %s %s\n", c.ClientIP(), c.Request.Method, c.Request.URL.Path))
		c.Next()
	}
}
