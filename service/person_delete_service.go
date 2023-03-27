package service

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func (service *personService) DeletePerson(personID int) error {

	var person model.Person
	err := service.PersonRepo.GetPersonWithPersonID(personID, &person)
	if err != nil {
		logg.Printlogger("DELETE failed, not found person id", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm delete failed, not found person id")
	}

	err = service.PersonRepo.DeletePerson(personID)
	if err != nil {
		logg.Printlogger("DELETE failed", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm delete failed")
	}

	return nil

}
