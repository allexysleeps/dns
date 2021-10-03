package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type DB interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

type redisDB struct {
	redis *redis.Client
}

var ctx = context.Background()

func (rdb redisDB) Set(key string, value string) error {
	err := rdb.redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return errors.New(fmt.Sprintf("unable to save: [%s] -> %v\n", key, value))
	}
	return nil
}

func (rdb redisDB) Get(key string) (string, error) {
	value, err := rdb.redis.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New(fmt.Sprintf("unable to get: [%s]\n", key))
	}
	return value, nil
}

func createRedisClient() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	
	return rdb
}

func Create() DB {
	rdb := createRedisClient()
	return redisDB{ redis: rdb }
}


