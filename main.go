package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var a App

var sqlOpen = sql.Open
var Println = fmt.Println

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/testtab", a.getTesttab).Methods("GET")
}

func (a *App) getTesttab(w http.ResponseWriter, r *http.Request) {
	dummyRequest := "xyz"
	respondWithJSON(w, http.StatusOK, dummyRequest)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Initialize(connectionString string) error {
	log.Printf("Connect to %v", connectionString)
	var err error
	a.DB, err = sqlOpen("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Printf("DB Status %v", a.DB.Stats())
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	return nil
}

func main() {
	Println(os.Getenv("TESTVAR"))
}
