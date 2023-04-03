package repository

import (
	"fmt"
	"time"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func (rc *redisCacheRepo) Update(newMackRedisCache model.MakeNewRedisCache) error {

	setKey := newMackRedisCache.Key

	if newMackRedisCache.Key != "" {

		expireTime, _ := time.ParseDuration(newMackRedisCache.Expire)
		err := rc.RedisCache.RedisClient.Set(rc.RedisCache.Contx, setKey, converse.JsonMarshal(newMackRedisCache.Data), expireTime).Err()
		if err != nil {
			logg.Printlogger("UPDATE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", setKey, err, logg.GetCallerPathNameFileNameLineNumber()))
			return err
		}

		return nil
	}

	logg.Printlogger("UPDATE Cache error", fmt.Sprintf("%v | %v | KEY =  %v", setKey, fmt.Errorf("can't update cahce :: key is empty or invalid format"), logg.GetCallerPathNameFileNameLineNumber()))
	return fmt.Errorf("can't Update cahce")
}
