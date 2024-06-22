package invokers

import (
	"context"
	"fmt"
	"go.uber.org/fx"

	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/pkg/process"
)

const ver = "20240620"

func RunRRDServer(process process.NameableProcessor, logger logger.ILogger, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go func() {
				logger.Info(fmt.Sprintf("starting process %s, ver: %s", process.Name(), ver))
				process.Run()
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info(fmt.Sprintf("stoping process %s, ver: %s", process.Name(), ver))
			process.Stop()
			return nil
		},
	})
}
