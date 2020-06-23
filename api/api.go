package api

import (
	"scratch/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logentry *logrus.Entry

func init() {
	logentry = log.Logger.WithField("api", "pongpong")
}

// Pong func
func Pong(c *gin.Context) {
	logentry.Info("call pong")
	c.String(200, "Pong")
}
