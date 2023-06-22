package cache

import (
	"fmt"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func RedisSettup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", setting.RedisSetting.Host, setting.RedisSetting.Port),
		Password: setting.RedisSetting.Password,
		DB:       0,
	})
}
