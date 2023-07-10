package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"

	"github.com/quan12xz/basic_japanese/pkg/setting"
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

func GenerateKey(pattern string, input ...any) string {
	key := fmt.Sprintf(pattern, input...)

	if key == "" {
		log.Fatal(errors.New("Key invalid"))
	}
	return key

}
