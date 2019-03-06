package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegration(t *testing.T) {
	a = App{}
	a.MountIntegration()
	err := a.Initialize()
	if err != nil {
		t.Errorf("Initialize not fail")
	}
	t.Run("Tables Exist", testTablesExist)
	t.Run("testTablesStore", testTablesStore)
	t.Run("testTablesStore", testGETtesttab)

}

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

func testTablesExist(t *testing.T) {
	tables := [...]string{"TESTTAB"}
	for _, table := range tables {
		_, e := a.dao.DB.Query("SELECT 1 FROM " + table + " LIMIT 1")
		if e != nil {
			t.Errorf("table %v does not exist %v", table, e)
		}
	}
}
func testTablesStore(t *testing.T) {
	query := "INSERT INTO TESTTAB (NAME) VALUES (NAME='Testname')"
	result, e := a.dao.DB.Exec(query)
	if e != nil {
		t.Errorf("Fehler beim bei \"%v\": %v", query, e)
	} else {
		i, _ := result.RowsAffected()
		if i != 1 {
			t.Errorf("Anzahl affected Rows nicht 1")
		}

	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.routing.router.ServeHTTP(rr, req)

	return rr
}
