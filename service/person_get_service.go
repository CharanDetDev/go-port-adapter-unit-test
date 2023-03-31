package service

import (
	"encoding/json"
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"gorm.io/gorm"
)

func (service *personService) GetPersonWithPersonID(personId int, person *model.Person) error {

	// * ----------------- Version Port & Adapter :: Get Redis cache -----------------
	key := fmt.Sprintf("person-%v", personId)
	personInCache, err := service.RedisCacheRepo.Get(key)
	if err == nil {
		personMarshal, _ := json.Marshal(personInCache)
		json.Unmarshal([]byte(personMarshal), &person)
		logg.Printlogger("********** GET Redis Cache Success **********", *person)
		return nil
	}
	// * -----------------------------------------------------------------------------

	err = service.PersonRepo.GetPersonWithPersonID(personId, person)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logg.PrintloggerJsonMarshalIndentHasHeader("********** GET failed **********", "Method GetPersonWithPersonID()", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
			return gorm.ErrRecordNotFound
		}
		return err
	}

	//* --------------------------- Set Redis cache ---------------------------------
	if person != nil {
		newMackRedisCache := model.MakeNewRedisCache{
			Key:    key,
			Data:   *person,
			Expire: "2m",
		}

		err = service.RedisCacheRepo.Set(newMackRedisCache)
		if err != nil {
			logg.PrintloggerJsonMarshalIndentHasHeader("********** SET Redis Cache failed **********", "Method GetPersonWithPersonID()", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
			return err
		}
	}
	// * -----------------------------------------------------------------------------

	return nil
}

// ? ----------------- Version ก่อนหน้า :: Get Redis cache -------------------------
// personInCache, err := cache.Get(fmt.Sprintf("%v", personId))
// if err == nil {
// 	personMarshal, _ := json.Marshal(personInCache)
// 	json.Unmarshal([]byte(personMarshal), &person)
// 	return nil
// }
// ? -----------------------------------------------------------------------------
