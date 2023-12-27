package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

type DatabaseUser struct {
	Db *redis.Client
}

func InitBaseUser() (*DatabaseUser, error) {

	redisBaseUser := os.Getenv("REDIS_USER_DATABASE")

	redisBaseUserInt, err := strconv.Atoi(redisBaseUser)
	if err != nil {
		return nil, err
	}

	client, err := AllDatabases.NewConnection(redisBaseUserInt)
	if err != nil {
		return nil, err
	}

	client.Set(context.Background(), "key", "value", 0)

	return &DatabaseUser{
		Db: client,
	}, nil

}
