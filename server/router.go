package server

import (
	"github.com/KiranGosavi/demoservice/server/handler"
	"github.com/gorilla/mux"
)

//creates a mux router, registers handler and returns router instance
func loadRoutes() *mux.Router {
	root := mux.NewRouter()
	root.StrictSlash(true)
	root.HandleFunc("/", handler.IndexHandler)
	root.HandleFunc("/website-details", handler.GetWebsiteDetails).Methods("GET")
	return root
}
