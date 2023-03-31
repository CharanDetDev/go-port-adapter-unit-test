package repository

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func (repo *personRepo) GetPersonWithPersonID(personId int, person *model.Person) error {

	err := repo.DatabaseConn.First(&person, personId).Error
	if err != nil {
		return err
	}

	return nil
}
