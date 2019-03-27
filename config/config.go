package config

import (
	"os"
	"strconv"

	newrelic "github.com/newrelic/go-agent"
)

type Config struct {
	logLevel       string
	host           string
	port           int
	mahakamIP      string
	mahakamPort    int
	hostBridgeName string
	newRelic       newrelic.Config
}

var appConfig *Config

func Load() error {
	appPort, err := strconv.Atoi(os.Getenv("app_port"))
	if err != nil {
		return err
	}

	mahakamPort, err := strconv.Atoi(os.Getenv("mahakam_port"))
	if err != nil {
		return err
	}

	nrConfig, err := getNewRelicConfigOrPanic()
	if err != nil {
		return err
	}

	// TODO(vjdhama): Do nill check for env vars
	appConfig = &Config{
		logLevel:       os.Getenv("log_level"),
		host:           os.Getenv("host"),
		port:           appPort,
		mahakamIP:      os.Getenv("mahakam_ip"),
		mahakamPort:    mahakamPort,
		hostBridgeName: os.Getenv("host_bridge_name"),
		newRelic:       nrConfig,
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

func MahakamIP() string {
	return appConfig.mahakamIP
}

func MahakamPort() int {
	return appConfig.mahakamPort
}

func HostBridgeName() string {
	return appConfig.hostBridgeName
}

func NewRelic() newrelic.Config {
	return appConfig.newRelic
}
