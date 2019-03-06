package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var a App

func main() {
	a.MountIntegration()
	a.Initialize()
	a.Println(os.Getenv("TESTVAR"))
}
