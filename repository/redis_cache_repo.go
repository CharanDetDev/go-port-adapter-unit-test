package repository

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/domain"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
)

type redisCacheRepo struct {
	database.RedisCache
}

func NewRedisCacheRepo(RedisCache database.RedisCache) domain.RedisCacheRepo {
	return &redisCacheRepo{
		RedisCache,
	}
}
