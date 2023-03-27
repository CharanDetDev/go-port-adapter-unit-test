package cache

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func Get(key string) (interface{}, error) {

	getByKey := key
	value, err := RedisCaching.RedisClient.Get(RedisCaching.Contx, getByKey).Result()
	if err != nil {

		if err.Error() == "redis: nil" {
			return nil, err
		}

		logg.Printlogger("GET Cache error", fmt.Sprintf("%v | %v | KEY =  %v", getByKey, err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return nil, err
	}

	return converse.JsonUnmarshal(value), nil
}
