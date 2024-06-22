package aerospike_ha_process

import (
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/pkg/process"

	procConfig "github.com/lissdx/aerospike-ha/internal/pkg/process/process_config"
	rrdService "github.com/lissdx/aerospike-ha/internal/service/http/rrd_service"
)

var _ process.NameableProcessor = (*AerospikeProcess)(nil)

type AerospikeProcess struct {
	process.Nameable
	logger     logger.ILogger
	rrdService *rrdService.RrdServer
	//echoServer *echo.Echo
}

func (asp *AerospikeProcess) Run() {

	asp.rrdService.Run()
	//err := asp.echoServer.Start("0.0.0.0:8080")
	//
	//if err != nil {
	//	asp.logger.Fatal(err.Error())
	//}
}

func (asp *AerospikeProcess) Stop() {
	asp.rrdService.Stop()
	//err := asp.echoServer.Shutdown(context.Background())
	//asp.rrdService.Driver.Stop()
	//if err != nil {
	//	asp.logger.Error(err)
	//}
}

func NewAerospikeProcess(pConf procConfig.ProcessConfigure, logger logger.ILogger, rrdService *rrdService.RrdServer) process.NameableProcessor {

	return &AerospikeProcess{
		Nameable:   process.NewNameable(pConf.GetProcessName()),
		logger:     logger,
		rrdService: rrdService,
	}
}
