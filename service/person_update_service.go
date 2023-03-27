package service

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func (service *personService) UpdatePerson(newPerson *model.PersonRequest) error {

	newPersonTemp := model.Person{
		PersonID:  newPerson.PersonID,
		LastName:  newPerson.LastName,
		FirstName: newPerson.FirstName,
		Address:   newPerson.Address,
		City:      newPerson.City,
	}

	var person model.Person
	err := service.PersonRepo.GetPersonWithPersonID(newPerson.PersonID, &person)
	if err != nil {
		logg.PrintloggerDebuggerHasHeader("********** UPDATE failed **********", "not found person id", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm update failed, not found person id")
	}

	err = service.PersonRepo.UpdatePerson(&newPersonTemp)
	if err != nil {
		logg.PrintloggerDebuggerHasHeader("********** UPDATE failed **********", "", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm update failed")
	}

	// fmt.Println("SUCCESS update by GORM")
	return nil

}
