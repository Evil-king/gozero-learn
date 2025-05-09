package db

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func NewRedisClient(redisConfig redis.RedisConf) *redis.Redis {
	return redis.MustNewRedis(redisConfig)
}
