package global

import (
	"github.com/golang/groupcache/singleflight"
	"github.com/kaijyin/md-server/server/config"
	"go.uber.org/zap"


	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MD_DB     *gorm.DB
	MD_REDIS  *redis.Client
	MD_CONFIG config.Server
	MD_Concurrency_Control = &singleflight.Group{}
	MD_VP     *viper.Viper
	MD_LOG    *zap.Logger
)
