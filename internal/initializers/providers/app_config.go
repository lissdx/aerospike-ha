package providers

import (
	strUtil "github.com/lissdx/aerospike-ha/internal/utils"
	"github.com/spf13/viper"
	"os"
)

const GoEnvEnvName = "GO_ENV"
const GoEnvDefault = "development"

func NewAppConfig() *viper.Viper {

	// set default GO_ENV if not existed
	env := strUtil.NormalizeStringToLower(os.Getenv(GoEnvEnvName))

	if strUtil.IsEmptyString(env) {
		env = GoEnvDefault
		err := os.Setenv(GoEnvEnvName, env)
		if err != nil {
			panic(err)
		}
	}

	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigName(env)
	config.AddConfigPath("./configs")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
