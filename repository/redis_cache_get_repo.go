package repository

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func (rc *redisCacheRepo) Get(key string) (interface{}, error) {

	getByKey := key
	value, err := rc.RedisCache.RedisClient.Get(rc.RedisCache.Contx, getByKey).Result()
	if err != nil {

		if err.Error() == "redis: nil" {
			return nil, err
		}

		logg.Printlogger("GET Cache error", fmt.Sprintf("%v | %v | KEY =  %v", getByKey, err, logg.GetCallerPathNameFileNameLineNumber()))
		return nil, err
	}

	return converse.JsonUnmarshal(value), nil
}
