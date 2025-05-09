package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	Auth        Auth
	RedisConfig redis.RedisConf
}

type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int64
}

type Auth struct {
	Secret string
	Expire int64
}

type RedisConfig struct {
	Host        string
	Type        string
	TLS         string
	NonBlock    bool
	PingTimeout int64
}
