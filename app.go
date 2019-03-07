package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/svenjeppsson/travis-test/interfaces"
	"log"
	"net/http"
	"strconv"
)

type Application interface {
	GetPerson(w http.ResponseWriter, r *http.Request)
	GetRouter() *mux.Router
}
type App struct {
	dal    interfaces.DataAccessLayer
	router *mux.Router
}

func (a *App) GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	_, person := a.dal.GetPerson(id)
	a.respondWithJSON(w, http.StatusOK, person)
}

func (a *App) GetRouter() *mux.Router {
	return a.router
}

func NewApp(dal interfaces.DataAccessLayer) Application {
	app := &App{
		dal:    dal,
		router: mux.NewRouter(),
	}
	app.router.HandleFunc("/person/{id}", app.GetPerson).Methods("GET")
	return app
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, e := w.Write(response)
	if e != nil {
		log.Printf("could write the response")
	}
}
