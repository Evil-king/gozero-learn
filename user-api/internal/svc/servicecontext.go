package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user-api/internal/config"
	"user-api/internal/db"
)

type ServiceContext struct {
	Config      config.Config
	MysqlCon    sqlx.SqlConn
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := db.NewMysql(c.MysqlConfig)
	RedisClient := db.NewRedisClient(c.RedisConfig)
	return &ServiceContext{
		Config:      c,
		MysqlCon:    con,
		RedisClient: RedisClient,
	}
}
