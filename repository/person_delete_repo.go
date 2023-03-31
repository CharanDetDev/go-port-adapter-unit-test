package repository

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func (repo *personRepo) DeletePerson(personID int) error {

	result := repo.DatabaseConn.Where("PersonID = ?", personID).Delete(&model.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
