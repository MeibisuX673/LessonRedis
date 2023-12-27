package database

import (
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

type DatabaseCache struct {
	Db *redis.Client
}

func InitBaseCache() (*DatabaseCache, error) {

	redisBaseCode := os.Getenv("REDIS_CACHE")

	redisBaseCodeInt, err := strconv.Atoi(redisBaseCode)
	if err != nil {
		return nil, err
	}

	client, err := AllDatabases.NewConnection(redisBaseCodeInt)
	if err != nil {
		return nil, err
	}

	return &DatabaseCache{
		Db: client,
	}, nil

}
