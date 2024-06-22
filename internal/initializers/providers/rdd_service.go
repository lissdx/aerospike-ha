package providers

import (
	"github.com/lissdx/aerospike-ha/internal/drivers/cache"
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/service/http/rrd_service"
	"github.com/spf13/viper"
)

func NewRrdService(config *viper.Viper, logger logger.ILogger, driver *cache.AerospikeDriver) *rrd_service.RrdServer {
	return rrd_service.NewRrdServer(config, logger, driver)
}
