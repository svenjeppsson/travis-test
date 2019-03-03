package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	t.Run("testTablesStore", testGETtesttab)

}

// main_test.go

func testGETtesttab(t *testing.T) {
	req, _ := http.NewRequest("GET", "/testtab", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := "\"xyz\""
	actual := response.Body.String()
	if actual != expected {
		t.Errorf("Expected %s Got %s", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

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
func TestApp_main(t *testing.T) {
	var test string
	Println = func(a ...interface{}) (n int, err error) {
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
