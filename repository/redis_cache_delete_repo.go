package repository

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func (rc *redisCacheRepo) Delete(keyTemp string) error {

	deleteKey := keyTemp
	err := rc.RedisCache.RedisClient.Del(rc.RedisCache.Contx, deleteKey).Err()
	if err != nil {
		logg.Printlogger("DELETE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", deleteKey, err, logg.GetCallerPathNameFileNameLineNumber()))
		return err
	}

	return nil
}
