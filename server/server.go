package server

import (
	"demoservice/config"
	"fmt"
	"net/http"
)

func Load()error{
	//load server config
	err :=config.Load()
	if err != nil {
		return err
	}
	return nil
}

func Start()error{
	router :=loadRoutes()
	return http.ListenAndServe(fmt.Sprintf(":%v",config.AppConfig.Port),router)
}