package repository

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func (repo *personRepo) GetPersonWithPersonID(personId int, person *model.Person) error {

	err := database.Conn.First(&person, personId).Error
	if err != nil {
		return err
	}

	return nil
}
