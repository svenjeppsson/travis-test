package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type App struct {
	dao     Dao
	routing Routing
	sqlOpen func(driverName string, dataSourceName string) (*sql.DB, error)
}

func (a *App) MountIntegration() {
	a.sqlOpen = sql.Open
}

func (a *App) Initialize() error {
	dbcon := os.Getenv("DBCON")
	err := a.dao.connect(dbcon)
	if err != nil {
		log.Printf("Could connext to %v: reason %v", dbcon, err)
		return err
	}
	a.routing.initializeRoutes()
	return nil
}

func (a *App) getTesttab(w http.ResponseWriter, r *http.Request) {
	dummyRequest := "xyz"
	a.respondWithJSON(w, http.StatusOK, dummyRequest)
}

//func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
//	a.respondWithJSON(w, code, map[string]string{"error": message})
//}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, e := w.Write(response)
	if e != nil {
		log.Printf("could write the response")
	}
}
