package app

func (server *Server) initializeRoute() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
