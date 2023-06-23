package cache

import (
	"errors"
	"fmt"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisSettup() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", setting.RedisSetting.Host, setting.RedisSetting.Port),
		Password: setting.RedisSetting.Password,
		DB:       0,
	})
}

func GenerateKey(input ...string) (string, error) {
	var key string
	for _, value := range input {
		if value == "" {
			return "", errors.New("invalid parameter to generate key")
		}

		key += value
	}
	return key, nil
}
