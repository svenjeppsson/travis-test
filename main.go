package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var a App

func main() {
	a.MountIntegration()
	err := a.Initialize()
	log.Fatal(err)
}
