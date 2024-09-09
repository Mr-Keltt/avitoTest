package shared

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Logger - global logger for the entire application.
var Logger = logrus.New()

// InitLogger - function for configuring the logger.
func InitLogger(config *Config) {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	level, err := logrus.ParseLevel(strings.ToLower(config.LogLevel))
	if err != nil {
		Logger.Warnf("Invalid log level '%s' provided. Falling back to 'info' level.", config.LogLevel)
		level = logrus.InfoLevel // Default level
	}
	Logger.SetLevel(level)

	Logger.SetOutput(os.Stdout)
}
