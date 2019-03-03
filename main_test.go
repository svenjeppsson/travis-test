package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	a = App{}
	err := a.Initialize(os.Getenv("DBCON"))
	if err != nil {
		t.Errorf("Initialize not fail")
	}
	t.Run("Tables Exist", testTablesExist)
	t.Run("testTablesStore", testTablesStore)

}

func TestApp_main(t *testing.T) {
	var test string
	Println = func(a ...interface{}) (n int, err error) {
		test = fmt.Sprint(a)
		return 0, nil
	}
	os.Setenv("TESTVAR", "testvalue")
	main()
	if test != "testvalue" {
		t.Error("main sould use TESTVAR")
	}
}

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

func testTablesExist(t *testing.T) {
	tables := [...]string{"TESTTAB"}
	for _, table := range tables {
		_, e := a.DB.Query("SELECT 1 FROM " + table + " LIMIT 1")
		if e != nil {
			t.Errorf("table %v does not exist %v", table, e)
		}
	}
}
func testTablesStore(t *testing.T) {
	query := "INSERT INTO TESTTAB (NAME) VALUES (NAME='Testname')"
	result, e := a.DB.Exec(query)
	if e != nil {
		t.Errorf("Fehler beim bei \"%v\": %v", query, e)
	} else {
		i, _ := result.RowsAffected()
		if i != 1 {
			t.Errorf("Anzahl affected Rows nicht 1")
		}

	}

}

func reInitExternalFunctions() {
	sqlOpen = sql.Open
}
