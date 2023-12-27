package database

import (
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

type DatabaseCode struct {
	Db *redis.Client
}

func InitBaseCode() (*DatabaseCode, error) {

	redisBaseCode := os.Getenv("REDIS_CODE_DATABASE")

	redisBaseCodeInt, err := strconv.Atoi(redisBaseCode)
	if err != nil {
		return nil, err
	}

	client, err := AllDatabases.NewConnection(redisBaseCodeInt)
	if err != nil {
		return nil, err
	}

	return &DatabaseCode{
		Db: client,
	}, nil

}
