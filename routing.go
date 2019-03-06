package main

import (
	"github.com/gorilla/mux"
)

type Routing struct {
	router *mux.Router
}

func (routing *Routing) initializeRoutes() {
	routing.router = mux.NewRouter()
	routing.router.HandleFunc("/testtab", a.getTesttab).Methods("GET")
}
