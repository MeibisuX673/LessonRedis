package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var AllDatabases Databases

type Databases struct {
	DbCode  *DatabaseCode
	DbUser  *DatabaseUser
	DbCache *DatabaseCache
}

func (db *Databases) Init() error {

	databaseCode, err := InitBaseCode()
	if err != nil {
		return err
	}
	db.DbCode = databaseCode

	dataBaseUser, err := InitBaseUser()
	if err != nil {
		return err
	}
	db.DbUser = dataBaseUser

	databaseCache, err := InitBaseCache()
	if err != nil {
		return err
	}
	db.DbCache = databaseCache

	return nil
}

func (db *Databases) NewConnection(base int) (*redis.Client, error) {

	redisUrl := os.Getenv("REDIS_DATABASE_URL")

	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       base,
	})
	status := client.Ping(context.Background())
	fmt.Println(status)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
