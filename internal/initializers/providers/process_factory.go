package providers

import (
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/pkg/process"
	procConfig "github.com/lissdx/aerospike-ha/internal/pkg/process/process_config"
	asHAProcess "github.com/lissdx/aerospike-ha/internal/pkg/process/processes/aerospike-ha-process"
	"github.com/lissdx/aerospike-ha/internal/service/http/rrd_service"
)

func ProcessorFactory(pConfig procConfig.ProcessConfigure, logger logger.ILogger, rrdService *rrd_service.RrdServer) process.NameableProcessor {
	switch pConfig.GetProcessName() {
	case "AEROSPIKE_HA":
		return asHAProcess.NewAerospikeProcess(pConfig, logger, rrdService)
	}

	logger.Panic("invalid process name")
	return nil
}
