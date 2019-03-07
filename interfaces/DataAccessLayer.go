package interfaces

import "github.com/svenjeppsson/travis-test/model"

type DataAccessLayer interface {
	StorePerson(person *model.Person) error
	DelelePerson(id int64) error
	GetPerson(id int64) (error, *model.Person)
	GetAllPersons() (error, []model.Person)
	GetPersonsBySearchString(search string) (error, []model.Person)
}
