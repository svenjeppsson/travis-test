package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/svenjeppsson/travis-test/dal"
)

var a App

func main() {
	NewApp(dal.NewSQLDataAcesssLayer())
	router := mux.NewRouter()
	router.HandleFunc("/persons", a.GetPerson).Methods("GET")

}
