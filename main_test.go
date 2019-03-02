package main

import (
	"database/sql"
	"errors"
	"testing"
)

func TestApp_Initialize(t *testing.T) {
	sqlOpen = func(driverName, dataSourceName string) (db *sql.DB, e error) {
		return nil, errors.New("Failed")
	}
	err := a.Initialize("ANY")
	if err == nil {
		t.Errorf("should fail wth error(\"ANY\")")
	}
	reInitExternalFunctions()
}

func reInitExternalFunctions() {
	sqlOpen = sql.Open
}

