package dal

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/svenjeppsson/travis-test/interfaces"
	"github.com/svenjeppsson/travis-test/model"
	"os"
	"strings"
)

type SQLDataAcesssLayer struct {
	DB *sqlx.DB
}

const NO_ROWS = "no rows"

func (dal *SQLDataAcesssLayer) StorePerson(person *model.Person) error {
	error := dal.connect()
	if error != nil {
		return error
	}
	result, error := dal.DB.NamedExec(`INSERT INTO PERSON(FIRSTNAME,LASTNAME) VALUES(:FIRSTNAME, :LASTNAME)`, person)

	if error != nil {
		return error
	}
	id, error := result.LastInsertId()

	if error != nil {
		return error
	}
	person.Id = &id
	return nil
}
func (dal *SQLDataAcesssLayer) DelelePerson(id int64) error {
	error := dal.connect()
	if error != nil {
		return error
	}
	result, error := dal.DB.Exec("DELETE FROM PERSON WHERE ID=?", id)

	const cantDelete = "Cant delete Person By Id=%v!"
	const cantDeleteWithReason = cantDelete + " Reason=%v"
	if error != nil {
		if strings.Contains(fmt.Sprintf("%v", error), NO_ROWS) {
			return fmt.Errorf(cantDeleteWithReason, id, error)
		} else {
			return error
		}

	}

	i, error := result.RowsAffected()
	if error != nil {
		return fmt.Errorf(cantDeleteWithReason, id, error)
	}

	if i < 1 {
		return fmt.Errorf(cantDelete, id)
	}

	return nil
}

func (dal *SQLDataAcesssLayer) GetPerson(id int64) (error, *model.Person) {
	error := dal.connect()
	if error != nil {
		return error, nil
	}
	person := &model.Person{}
	error = dal.DB.Get(person, "SELECT ID,LASTNAME,FIRSTNAME FROM PERSON WHERE ID=?", id)
	if error != nil {
		if strings.Contains(fmt.Sprintf("%v", error), "no rows") {
			return fmt.Errorf("Cant find Person By Id=%v", id), nil
		}
		return error, nil
	}
	return nil, person
}

func (dal *SQLDataAcesssLayer) GetAllPersons() (error, []model.Person) {
	error := dal.connect()
	if error != nil {
		return error, nil
	}
	persons := []model.Person{}
	error = dal.DB.Select(&persons, "SELECT * FROM PERSON")
	return error, persons
}

func (dal *SQLDataAcesssLayer) GetPersonsBySearchString(search string) (error, []model.Person) {
	error := dal.connect()
	if error != nil {
		return error, nil
	}
	persons := []model.Person{}
	error = dal.DB.Select(&persons, "SELECT * FROM PERSON WHERE LOWER(CONCAT(FIRSTNAME,LASTNAME)) LIKE ?", "%"+strings.ToLower(search)+"%")
	return error, persons
}

func NewSQLDataAcesssLayer() interfaces.DataAccessLayer {
	sqlDataAcesssLayer := SQLDataAcesssLayer{}
	return &sqlDataAcesssLayer
}

func (dal *SQLDataAcesssLayer) connect() error {
	var err error = nil
	if dal.DB == nil {
		dal.DB, err = sqlx.Open("mysql", os.Getenv("DBCON"))
	}
	return err
}
