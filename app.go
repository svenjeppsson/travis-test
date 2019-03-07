package main

import (
	"encoding/json"
	"github.com/svenjeppsson/travis-test/interfaces"
	"log"
	"net/http"
)

type Application interface {
	GetPerson(w http.ResponseWriter, r *http.Request)
}
type App struct {
	dal interfaces.DataAccessLayer
}

func (a *App) GetPerson(w http.ResponseWriter, r *http.Request) {
	dummyRequest := "xyz"
	a.respondWithJSON(w, http.StatusOK, dummyRequest)
}

func NewApp(dal interfaces.DataAccessLayer) Application {
	app := &App{
		dal: dal,
	}
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
