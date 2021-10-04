package redisStorage

import (
  "context"
  "errors"
  "fmt"
  "github.com/go-redis/redis/v8"
  "log"
)

type RedisDB struct {
  redis *redis.Client
}

var ctx = context.Background()

func (rdb RedisDB) Set(key string, value string) error {
  err := rdb.redis.Set(ctx, key, value, 0).Err()
  if err != nil {
    log.Print(err)
    return errors.New(fmt.Sprintf("unable to save: [%s] -> %v\n", key, value))
  }
  return nil
}

func (rdb RedisDB) Get(key string) (string, error) {
  value, err := rdb.redis.Get(ctx, key).Result()
  if err != nil {
    return "", errors.New(fmt.Sprintf("unable to get: [%s]\n", key))
  }
  return value, nil
}

func  CreateRedisClient() RedisDB {
  rdb := redis.NewClient(&redis.Options{
    Addr: "redis:6379",
    Password: "",
    DB: 0,
  })
  
  return RedisDB{ redis: rdb }
}