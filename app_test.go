package main

import (
	"encoding/json"
	"github.com/svenjeppsson/travis-test/mocks"
	"github.com/svenjeppsson/travis-test/model"
	"github.com/svenjeppsson/travis-test/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app Application

func TestApp(t *testing.T) {
	layer := mocks.DataAccessLayer{}
	layer.On("GetPerson", int64(1)).Return(nil, &model.Person{Id: util.AdrInt64(1), FirstName: "Max", LastName: "Maier"})
	app = NewApp(&layer)
	testGETPerson(t)
}

// main_test.go

func testGETPerson(t *testing.T) {
	req, _ := http.NewRequest("GET", "/person/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	bytes, _ := json.Marshal(model.Person{Id: util.AdrInt64(1), FirstName: "Max", LastName: "Maier"})
	actual := response.Body.String()
	if actual != string(bytes) {
		t.Errorf("Expected %s Got %s", string(bytes), actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.GetRouter().ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
