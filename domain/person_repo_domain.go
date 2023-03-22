package domain

import "github.com/CharanDetDev/go-port-adapter-unit-test/model"

type PersonRepo interface {
	GetPersonWithPersonID(personId int) (*model.Person, error)
	InsertPerson(newPerson *model.Person) error
	UpdatePerson(newPerson *model.Person) error
	DeletePerson(personID int) error
}
