package dal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/svenjeppsson/travis-test/interfaces"
	"github.com/svenjeppsson/travis-test/model"
	"github.com/svenjeppsson/travis-test/util"
	"log"
	"os"
	"testing"
)

var sqldal interfaces.DataAccessLayer

func TestIntegrationDataAccessLayer(t *testing.T) {
	sqldal = NewSQLDataAcesssLayer()
	truncateTable("PERSON")
	testSQLDataAcesssLayer_GetAllPersons(0, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("max", 0, t)
	testSQLDataAcesssLayer_GetPerson(1, nil, fmt.Errorf("Cant find Person By Id=%v", 1), t)
	testSQLDataAcesssLayer_DeletePerson(0, fmt.Errorf("Cant delete Person By Id=%v", 0), t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max", LastName: "Meier"}, &model.Person{Id: util.AdrInt64(1), FirstName: "Max", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max", LastName: "Meier"}, &model.Person{FirstName: "Max", LastName: "Meier"}, util.AdrStr("Duplicate"), t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max2", LastName: "Meier"}, &model.Person{Id: util.AdrInt64(3), FirstName: "Max2", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_GetAllPersons(2, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("Max", 2, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("axmei", 1, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("nix", 0, t)
	testSQLDataAcesssLayer_GetPerson(1, &model.Person{Id: util.AdrInt64(1), FirstName: "Max", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_DeletePerson(1, nil, t)

}

func testSQLDataAcesssLayer_StorePerson(person *model.Person, expectedPerson *model.Person, expectederror *string, t *testing.T) {
	error := sqldal.StorePerson(person)
	if expectederror != nil {
		assert.Contains(t, fmt.Sprintf("%v", error), *expectederror)
	} else {
		assert.Nil(t, error)
	}
	assert.Equal(t, expectedPerson, person)

}
func testSQLDataAcesssLayer_GetPerson(id int64, expectedPerson *model.Person, expectederror error, t *testing.T) {
	error, person := sqldal.GetPerson(id)
	assert.Equal(t, expectederror, error)
	assert.Equal(t, expectedPerson, person)

}
func testSQLDataAcesssLayer_DeletePerson(id int64, expectederror error, t *testing.T) {
	error := sqldal.DelelePerson(id)
	assert.Equal(t, expectederror, error)

}
func testSQLDataAcesssLayer_GetPersonsBySearchString(search string, expectedCount int, t *testing.T) {
	error, persons := sqldal.GetPersonsBySearchString(search)
	if error != nil {
		t.Errorf("%v", error)
	}
	assert.Equal(t, expectedCount, len(persons))
}

func testSQLDataAcesssLayer_GetAllPersons(expectedCount int, t *testing.T) {
	error, persons := sqldal.GetAllPersons()
	if error != nil {
		t.Error(error)
	}
	assert.Equal(t, expectedCount, len(persons))
}

func truncateTable(tablename string) {
	dbcon := os.Getenv("DBCON")
	db, err := sqlx.Connect("mysql", dbcon)
	if err != nil {
		log.Fatalf("error by connection to %v : %v", dbcon, err)
	}
	db.MustExec("TRUNCATE TABLE " + tablename)
	db.Close()
}
