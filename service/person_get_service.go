package service

import (
	"encoding/json"
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/cache"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func (service *personService) GetPersonWithPersonID(personId int) (*model.Person, error) {

	var person *model.Person

	// * :: ====================================================
	// * :: Get Redis cache
	personInCache, err := cache.Get(fmt.Sprintf("%v", personId))
	if err == nil {
		personMarshal, _ := json.Marshal(personInCache)
		json.Unmarshal([]byte(personMarshal), &person)
		return person, nil
	}
	// * :: ====================================================

	person, err = service.PersonRepo.GetPersonWithPersonID(personId)
	if err != nil {
		logg.Printlogger("GET failed", "", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return nil, fmt.Errorf("gorm personID not found")
	}

	return person, nil

}
