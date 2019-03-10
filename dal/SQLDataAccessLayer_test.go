package dal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/svenjeppsson/travis-test/interfaces"
	"github.com/svenjeppsson/travis-test/model"
	"log"
	"os"
	"testing"
)

var sqldal interfaces.DataAccessLayer

func TestDataAccessLayer(t *testing.T) {
	sqldal = NewSQLDataAcesssLayer()

	truncateTable("PERSON")
	testSQLDataAcesssLayer_GetAllPersons(0, nil, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("max", 0, nil, t)
	testSQLDataAcesssLayer_GetPerson(1, nil, AdrStr(fmt.Sprintf("Cant find Person By Id=%v", 1)), t)
	testSQLDataAcesssLayer_DeletePerson(0, AdrStr(fmt.Sprintf("Cant delete Person By Id=%v", 0)), t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max", LastName: "Meier"}, &model.Person{Id: AdrInt64(1), FirstName: "Max", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max", LastName: "Meier"}, &model.Person{FirstName: "Max", LastName: "Meier"}, AdrStr("Duplicate"), t)
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max2", LastName: "Meier"}, &model.Person{Id: AdrInt64(3), FirstName: "Max2", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_GetAllPersons(2, nil, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("Max", 2, nil, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("axmei", 1, nil, t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("nix", 0, nil, t)
	testSQLDataAcesssLayer_GetPerson(1, &model.Person{Id: AdrInt64(1), FirstName: "Max", LastName: "Meier"}, nil, t)
	testSQLDataAcesssLayer_DeletePerson(1, nil, t)

}

func TestDataAccessLayerConnectError(t *testing.T) {
	os.Setenv("DBCON", "dings")
	sqldal = NewSQLDataAcesssLayer()
	testSQLDataAcesssLayer_StorePerson(&model.Person{FirstName: "Max", LastName: "Meier"}, &model.Person{FirstName: "Max", LastName: "Meier"}, AdrStr("invalid DSN"), t)
	testSQLDataAcesssLayer_GetPerson(100, nil, AdrStr("invalid DSN"), t)
	testSQLDataAcesssLayer_GetPersonsBySearchString("Max", 0, AdrStr("invalid DSN"), t)
	testSQLDataAcesssLayer_GetAllPersons(0, AdrStr("invalid DSN"), t)
	testSQLDataAcesssLayer_DeletePerson(17, AdrStr("invalid DSN"), t)
}

func testSQLDataAcesssLayer_StorePerson(person *model.Person, expectedPerson *model.Person, expectederror *string, t *testing.T) {
	error := sqldal.StorePerson(person)
	expectError(expectederror, t, error)
	assert.Equal(t, expectedPerson, person)

}

func testSQLDataAcesssLayer_GetPerson(id int64, expectedPerson *model.Person, expectederror *string, t *testing.T) {
	error, person := sqldal.GetPerson(id)
	expectError(expectederror, t, error)
	assert.Equal(t, expectedPerson, person)

}
func testSQLDataAcesssLayer_DeletePerson(id int64, expectederror *string, t *testing.T) {
	error := sqldal.DelelePerson(id)
	expectError(expectederror, t, error)

}
func testSQLDataAcesssLayer_GetPersonsBySearchString(search string, expectedCount int, expectederror *string, t *testing.T) {
	error, persons := sqldal.GetPersonsBySearchString(search)
	expectError(expectederror, t, error)
	assert.Equal(t, expectedCount, len(persons))
}

func testSQLDataAcesssLayer_GetAllPersons(expectedCount int, expectederror *string, t *testing.T) {
	error, persons := sqldal.GetAllPersons()
	expectError(expectederror, t, error)
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

func expectError(expectederror *string, t *testing.T, e error) {
	if expectederror != nil {
		assert.Contains(t, fmt.Sprintf("%v", e), *expectederror)
	} else {
		assert.Nil(t, e)
	}
}

func AdrInt64(v int64) *int64 {
	return &v
}
func AdrStr(v string) *string {
	return &v
}
