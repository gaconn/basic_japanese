package cache

import (
	"context"
	"encoding/json"
	"time"
)

type RedisSetup struct {
	Context    context.Context
	Key        string
	Value      interface{}
	ExpireTime time.Duration
}

func (r *RedisSetup) GetData(object interface{}) error {
	res, err := RedisClient.Get(r.Context, r.Key).Result()

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
	_, err := RedisClient.Set(r.Context, r.Key, r.Value, r.ExpireTime).Result()

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSetup) DeleteData() error {
	_, err := RedisClient.Del(r.Context, r.Key).Result()

	if err != nil {
		return err
	}
	return nil
}
