package service

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"gorm.io/gorm"
)

func (service *personService) GetPersonWithPersonID(personId int, person *model.Person) error {

	err := service.PersonRepo.GetPersonWithPersonID(personId, person)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logg.PrintloggerJsonMarshalIndentHasHeader("********** GET failed **********", "Method GetPersonWithPersonID()", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

// * :: ====================================================
// * :: Get Redis cache
// personInCache, err := cache.Get(fmt.Sprintf("%v", personId))
// if err == nil {
// 	personMarshal, _ := json.Marshal(personInCache)
// 	json.Unmarshal([]byte(personMarshal), &person)
// 	return person, nil
// }
// * :: ====================================================
