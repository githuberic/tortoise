package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RedisDb *redis.Client
)

func SetupRedisDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	RedisDb = redis.NewClient(&redis.Options{
		Addr:         RedisSetting.Addr,
		Password:     RedisSetting.Password, // no password set
		DB:           RedisSetting.DB,       // use default DB
		DialTimeout:  time.Millisecond * RedisSetting.DialConnectionTimeout,
		ReadTimeout:  time.Millisecond * RedisSetting.DialReadTimeout,
		WriteTimeout: time.Millisecond * RedisSetting.DialWriteTimeout,
		IdleTimeout:  time.Millisecond * RedisSetting.IdleTimeout,
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
