package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//You can make every method mockable by creating a func
//In main.go we do
//	var Println = fmt.Println
//In the the testmethod we can now moch Println
//	var test string
//	Println = func(a ...interface{}) (n int, err error) {
//		test = fmt.Sprintf("%v",a[0])
//		return 0, nil
//	}
func Test_main(t *testing.T) {
	var test string
	a.Println = func(a ...interface{}) (n int, err error) {
		test = fmt.Sprintf("%v", a[0])
		return 0, nil
	}
	os.Setenv("TESTVAR", "testvalue")
	main()
	if test != "testvalue" {
		t.Errorf("main should print the content of the env var TESTVAR: '%v' but it prints '%v' ", os.Getenv("TESTVAR"), test)
	}
}

//To Test the error handling you can force a mocked Method to return an error
//	sqlOpen = func(driverName, dataSourceName string) (db *sql.DB, e error) {
//		return nil, errors.New("Failed")
//	}
func TestApp_Initialize(t *testing.T) {
	a.sqlOpen = func(driverName, dataSourceName string) (db *sql.DB, e error) {
		return nil, errors.New("Failed")
	}
	err := a.Initialize()
	if err != nil {
		t.Errorf("should fail with error(Failed)")
	}
	a.MountIntegration()
}
