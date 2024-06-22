package providers

import (
	procConfig "github.com/lissdx/aerospike-ha/internal/pkg/process/process_config"
	"github.com/lissdx/aerospike-ha/internal/utils"
	"github.com/spf13/viper"
)

const ProcessNameEnv = "PROCESS_NAME"
const DefaultProcessName = "AEROSPIKE_HA"

func NewProcessConfig(config *viper.Viper) procConfig.ProcessConfigure {

	procName := utils.StrOrDefault(config.GetString(ProcessNameEnv), DefaultProcessName)

	return procConfig.NewProcessConfigure(procName)
}
