package model

type Person struct {
	Id        *int64 `db:"ID" json:"id"`
	FirstName string `db:"FIRSTNAME" json:"firstname"`
	LastName  string `db:"LASTNAME" json:"lastname"`
}
