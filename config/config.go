package config

import (
	"encoding/json"
	"os"
)

//structure to hold configuration parameters for the demoservice
type appConfig struct {
	Port     int    `json:"demo_service_port"`
	LogLevel string `json:"log_level"`
}

//global variable
var AppConfig appConfig

func Load() error {
	appConfig, present := os.LookupEnv("DEMO_SERVER_PROPERTIES")
	if present {
		err := json.Unmarshal([]byte(appConfig), &AppConfig)
		if err != nil {
			return err
		}
	} else {
		AppConfig.Port = 8080
		AppConfig.LogLevel = "DEBUG"
	}
	return nil
}
