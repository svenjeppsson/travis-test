package main

import (
	"database/sql"
	"errors"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
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
	if err == nil {
		t.Errorf("should fail with error(Failed)")
	}
	a.MountIntegration()
}
