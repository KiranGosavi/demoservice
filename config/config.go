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
	appconfig := os.Getenv("DEMO_SERVER_PROPERTIES")
	err := json.Unmarshal([]byte(appconfig), &AppConfig)
	if err != nil {
		return err
	}
	return nil
}
