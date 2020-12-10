package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"master/resource"
)

var (
	RedisDB *redis.Client
	RedisDBCtx context.Context
)

func init() {
	cfg, _ := ini.Load(resource.ConfFilePath)

	redisCfg := cfg.Section("redis")
	address := redisCfg.Key("host").String() + ":" + redisCfg.Key("port").String()
	db, _ := redisCfg.Key("db").Int()

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       db,
	})
	RedisDBCtx = context.Background()
}
