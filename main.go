package main

import (
	"scratch/api"
	"scratch/log"
	"scratch/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()
	r.Use(middleware.MyLogger(), gin.Recovery())
	/* r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%v | %v | %v | %v |", params.ClientIP, params.Latency, params.Method, params.Request.URL)
	})) */

	r.GET("/", func(c *gin.Context) {
		log.Logger.WithFields(logrus.Fields{
			"name": "hah",
		}).Info("call index")

		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	r.GET("/ping", api.Pong)

	r.Run()
}
