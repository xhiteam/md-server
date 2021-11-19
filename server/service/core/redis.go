package core

import (
	"context"
	"github.com/kaijyin/md-server/server/global"
	"time"
)

type RedisService struct {
}

var ctx = context.Background()
var RedisServiceApp = new(RedisService)

func (redisService *RedisService) Delete(userName string) (err error) {
	err = global.MD_REDIS.Del(ctx,userName).Err()
	return err
}

func (redisService *RedisService) Get(key string) (err error, val string) {
	val, err = global.MD_REDIS.Get(ctx, key).Result()
	return err, val
}

func (redisService *RedisService) Set(key string, val string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.MD_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.MD_REDIS.Set(ctx, val, key, timer).Err()
	return err
}
