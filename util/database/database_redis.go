package database

import (
	"context"
	"strconv"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/redis/go-redis/v9"
)

var RedisCaching RedisCache

type RedisCache struct {
	RedisClient *redis.Client   `json:"redis_client"`
	EnvPrefix   string          `json:"env_prefix"`
	Contx       context.Context `json:"contx"`
}

func InitRedisCache() bool {

	redisDBNum, _ := strconv.Atoi(config.Env.REDIS_DB_NUM)
	RedisCaching.Contx = context.Background()
	RedisCaching.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Env.REDIS_ADDRESS,
		Password: config.Env.REDIS_PASSWORD,
		DB:       redisDBNum,
	})

	return true
}

func RedisCacheClose() {
	RedisCaching.RedisClient.Close()
}
