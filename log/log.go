package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger logger
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
	f, _ := os.Create("gin.log")
	Logger.SetOutput(f)
}
