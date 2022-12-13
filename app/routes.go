package app

import (
	"github.com/gorilla/mux"
	"github.com/taujago/gotoko/app/controllers"
)

func (server *Server) initializeRoute() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
