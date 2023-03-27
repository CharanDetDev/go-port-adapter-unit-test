package cache

import (
	"fmt"
	"time"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func Update(newCache MakeCache) error {

	setKey := newCache.Key

	if newCache.Key != "" {

		expireTime, _ := time.ParseDuration(newCache.Expire)
		err := RedisCaching.RedisClient.Set(RedisCaching.Contx, setKey, converse.JsonMarshal(newCache.Data), expireTime).Err()
		if err != nil {
			return err
		}

		logg.Printlogger("UPDATE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", setKey, err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return nil
	}

	logg.Printlogger("UPDATE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", setKey, fmt.Errorf("can't update cahce :: key is empty or invalid format"), logg.GetCallerPathNameFileNameLineNumber()))
	return fmt.Errorf("can't set cahce")
}
