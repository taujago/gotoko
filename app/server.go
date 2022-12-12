package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Welcome to toko")
	server.Router = mux.NewRouter()
}

func (server *Server) Run(addr string) {
	fmt.Println(".. Listenin on port ,", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	server.Initialize()
	server.Run("9000")
}
