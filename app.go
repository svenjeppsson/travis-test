package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type App struct {
	dao     Dao
	routing Routing
	sqlOpen func(driverName string, dataSourceName string) (*sql.DB, error)
	Println func(a ...interface{}) (n int, err error)
}

func (a *App) MountIntegration() {
	a.sqlOpen = sql.Open
	a.Println = fmt.Println
}

func (a *App) Initialize() error {
	err := a.dao.connect(os.Getenv("DBCON"))
	if err != nil {
		return err
	}
	a.routing.initializeRoutes()
}

func (a *App) getTesttab(w http.ResponseWriter, r *http.Request) {
	dummyRequest := "xyz"
	a.respondWithJSON(w, http.StatusOK, dummyRequest)
}

func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
