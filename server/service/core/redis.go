package core

import (
	"context"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/table"
	"time"
)

type RedisService struct {
}
var RedisServiceApp=new(RedisService)
func (redisService *RedisService) CheckAuthority(userId uint,contextId uint,permission table.PermissionType)error{

}



func (redisService *RedisService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.MD_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}


func (redisService *RedisService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.MD_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.MD_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

