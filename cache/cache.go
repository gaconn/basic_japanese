package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/quan12xz/basic_japanese/pkg/setting"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	RedisClient *redis.Client
	Context     context.Context
}

var redisClientInstance *redisClient

func GetInstance() *redisClient {
	if redisClientInstance == nil {
		redisClientInstance = &redisClient{
			RedisClient: redisSettup(),
			Context:     context.TODO(),
		}
	}
	return redisClientInstance
}
func redisSettup() *redis.Client {
	var redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", setting.RedisSetting.Host, setting.RedisSetting.Port),
		Password: setting.RedisSetting.Password,
		DB:       0,
	})
	return redisClient
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
