package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient() *RedisClient {
	rd := &RedisClient{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}

	return rd
}

func (rd *RedisClient) RedisSet(key, value string) {
	err := rd.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (rd *RedisClient) RedisGet(key string) (string, error) {
	val, err := rd.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
