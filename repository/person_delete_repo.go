package repository

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func (repo *personRepo) DeletePerson(personID int) error {

	result := database.Conn.Where("PersonID = ?", personID).Delete(&model.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
