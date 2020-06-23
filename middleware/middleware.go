package middleware

import (
	"scratch/log"

	"github.com/gin-gonic/gin"
)

// MyLogger func return logger mw
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.Infof("%v | %v | %v |", c.ClientIP(), c.Request.Method, c.Request.URL)
		c.Next()
	}
}
