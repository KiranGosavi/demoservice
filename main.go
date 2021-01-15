package main

import (
	"github.com/KiranGosavi/demoservice/server"
	"fmt"
)

func main() {
	fmt.Println("Loading server configuration...")
	err := server.Load()
	if err != nil {
		panic(fmt.Sprint("Error while loading server configuration: %v", err))
	}
	err = server.Start()
	if err != nil {
		panic(fmt.Sprintf("Error while starting server : %v", err))
	}

}
