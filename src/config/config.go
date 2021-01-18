package config

import (
	"github.com/kelseyhightower/envconfig"
)

var runtimeConfig Config

type Config struct {
	ProcessName string
	LogDir      string
	Port        int
}

func Get() Config {
	return runtimeConfig
}

func Setup(processName string) error {
	err := envconfig.Process(processName, &runtimeConfig)
	if err != nil {
		return err
	}

	runtimeConfig.ProcessName = processName

	return nil
}
