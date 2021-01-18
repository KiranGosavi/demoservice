package server

import (
	"fmt"
	"github.com/KiranGosavi/demoservice/config"
	"net/http"
)

func Load() error {
	//load server config
	err := config.Load()
	if err != nil {
		return err
	}
	return nil
}

func Start() error {
	router := loadRoutes()
	//start web server on a given port
	return http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router)
}
