package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Serve() {
	// Start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// projectService := NewProjectService(s.store)
	// projectService.RegisterRoutes(subrouter)

	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subRouter)

	// registering our services
	log.Println("Starting the API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subRouter))

}
