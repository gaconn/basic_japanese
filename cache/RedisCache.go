package cache

import (
	"encoding/json"
	"time"
)

type RedisSetup struct {
	Key        string
	Value      interface{}
	ExpireTime time.Duration
}

func (r *RedisSetup) GetData(object interface{}) error {
	redisClientInstance := GetInstance()
	res, err := redisClientInstance.RedisClient.Get(redisClientInstance.Context, r.Key).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(res), object)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSetup) SetData() error {
	redisClientInstance := GetInstance()
	_, err := redisClientInstance.RedisClient.Set(redisClientInstance.Context, r.Key, r.Value, r.ExpireTime).Result()

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSetup) DeleteData() error {
	redisClientInstance := GetInstance()
	_, err := redisClientInstance.RedisClient.Del(redisClientInstance.Context, r.Key).Result()

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSetup) GetDataTest() string {
	redisClientInstance := GetInstance()
	result, err := redisClientInstance.RedisClient.Get(redisClientInstance.Context, r.Key).Result()

	if err != nil {
		return err.Error()
	}
	return result
}
