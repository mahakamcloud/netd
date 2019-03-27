package config

import (
	"os"
	"testing"
)

func Test_ConfigIsLoadedProperly(t *testing.T) {
	configVars := map[string]string{
		"log_level":             "debug",
		"host":                  "localhost",
		"app_port":              "3000",
		"new_relic_enabled":     "false",
		"mahakam_ip":            "1.2.3.4",
		"mahakam_port":          "9001",
		"new_relic_app_name":    "dummy",
		"host_bridge_name":      "mbr0",
		"new_relic_license_key": "dummy_license_key",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	Load()

	if appConfig.logLevel != "debug" {
		t.Errorf("Wrong log level. Got: %v Want: %v", appConfig.logLevel, "debug")
	}
	if appConfig.port != 3000 {
		t.Errorf("Wrong app port. Got: %v Want: %v", appConfig.port, 3000)
	}
	if appConfig.host != "localhost" {
		t.Errorf("Wrong host. Got: %v Want: %v", appConfig.host, "localhost")
	}
	if appConfig.hostBridgeName != "mbr0" {
		t.Errorf("Wrong host bridg name. Got: %v Want: %v", appConfig.hostBridgeName, "mbr0")
	}
	if appConfig.mahakamIP != "1.2.3.4" {
		t.Errorf("Wrong mahakam IP. Got: %v Want: %v", appConfig.mahakamIP, "1.2.3.4")
	}
	if appConfig.mahakamPort != 9001 {
		t.Errorf("Wrong mahakam port. Got: %v Want: %v", appConfig.mahakamPort, 9001)
	}
	if appConfig.newRelic.Enabled != false {
		t.Errorf("Wrong New Relic enabled. Got: %v Want: %v", appConfig.newRelic.Enabled, false)
	}
	if appConfig.newRelic.AppName != "dummy" {
		t.Errorf("Wrong New Relic app name. Got: %v Want: %v", appConfig.newRelic.AppName, "dummy")
	}
	if appConfig.newRelic.License != "dummy_license_key" {
		t.Errorf("Wrong New Relic license key. Got: %v Want: %v", appConfig.newRelic.License, "dummy_license_key")
	}
}

func Test_LoadConfigShouldFailIfPortIsNotInteger(t *testing.T) {
	configVars := map[string]string{
		"app_port": "port",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	err := Load()
	if err == nil {
		t.Errorf("Config load should have failed for value: port")
	}
}

func Test_LoadConfigShouldFailIfNewRelicEnabledNotBool(t *testing.T) {
	configVars := map[string]string{
		"app_port":          "3000",
		"new_relic_enabled": "foobar",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	err := Load()
	if err == nil {
		t.Errorf("Loading New Relic config should have failed")
	}
}
