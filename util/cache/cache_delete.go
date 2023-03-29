package cache

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func Delete(keyTemp string) error {

	deleteKey := keyTemp
	err := RedisCaching.RedisClient.Del(RedisCaching.Contx, deleteKey).Err()
	if err != nil {
		logg.Printlogger("DELETE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", deleteKey, err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return err
	}

	logg.Printlogger("DELETE Cache error", fmt.Sprintf(" %v | KEY =  %v", deleteKey, logg.GetCallerPathNameFileNameLineNumber()))
	return nil
}