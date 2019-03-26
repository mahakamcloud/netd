package config

import (
	"os"
	"strconv"

	newrelic "github.com/newrelic/go-agent"
)

type Config struct {
	logLevel string
	host     string
	port     int
	newRelic newrelic.Config
}

var appConfig *Config

func Load() error {
	appPort, err := strconv.Atoi(os.Getenv("app_port"))
	if err != nil {
		return err
	}

	nrConfig, err := getNewRelicConfigOrPanic()
	if err != nil {
		return err
	}

	appConfig = &Config{
		logLevel: os.Getenv("log_level"),
		host:     os.Getenv("host"),
		port:     appPort,
		newRelic: nrConfig,
	}
	return nil
}

func getNewRelicConfigOrPanic() (newrelic.Config, error) {
	nrEnabled, err := strconv.ParseBool(os.Getenv("new_relic_enabled"))
	if err != nil {
		return newrelic.Config{}, err
	}
	nrConfig := newrelic.NewConfig(os.Getenv("new_relic_app_name"), os.Getenv("new_relic_license_key"))
	nrConfig.Enabled = nrEnabled
	return nrConfig, nil
}

func Port() int {
	return appConfig.port
}

func NewRelic() newrelic.Config {
	return appConfig.newRelic
}
