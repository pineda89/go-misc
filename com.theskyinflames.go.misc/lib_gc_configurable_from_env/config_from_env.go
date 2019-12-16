package lib_gc_configurable_from_env

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

const (
	ENV_VARIABLES_PREFIX = "VAR_PREFIX"
)

func init() {
}

func GetEnvVariable(name string) (string, error) {

	if value := viper.Get(name); value == nil {
		return "", errors.New(fmt.Sprintf("Environment variable %s is not defined !!!", name))
	} else {
		return value.(string), nil
	}
}

func LoadVariablesFromEnv(prefix string) error {

	viper.SetEnvPrefix(prefix) // will be uppercased automatically
	viper.AutomaticEnv()

	return nil
}
