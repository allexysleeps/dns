package storage

import (
	"github.com/allexysleeps/dns/dns-storage/storage/redisStorage"
)

type DB interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

func Create() DB {
	rdb := redisStorage.CreateRedisClient()
	return rdb
}


