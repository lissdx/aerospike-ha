package providers

import (
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) logger.ILogger {
	return logger.LoggerFactory(logger.WithZapLoggerImplementer(),
		logger.WithLoggerLevel(config.GetString("LOG_LEVEL")))
}
