package domain

import "github.com/CharanDetDev/go-port-adapter-unit-test/model"

type RedisCacheRepo interface {
	Delete(keyTemp string) error
	Get(key string) (interface{}, error)
	Set(newMackRedisCache model.MakeNewRedisCache) error
	Update(newMackRedisCache model.MakeNewRedisCache) error
}
