package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var a App

var sqlOpen = sql.Open

type App struct {
	DB *sql.DB
}

func (a *App) Initialize(connectionString string) error {
	log.Printf("Connect to %v", connectionString)
	var err error
	a.DB, err = sqlOpen("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Printf("DB Status %v", a.DB.Stats())
	return nil
}

func main() {

}
