package providers

import (
	"github.com/lissdx/aerospike-ha/internal/drivers/cache"
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/spf13/viper"
)

func NewAerospikeDriver(config *viper.Viper, logger logger.ILogger) *cache.AerospikeDriver {
	return cache.NewAerospikeDriver(config, logger)
}
