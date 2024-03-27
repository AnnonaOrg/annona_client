package dbredis

import (
	"context"
	"time"

	"github.com/AnnonaOrg/annona_client/internal/repository"

	"github.com/AnnonaOrg/osenv"
	"github.com/redis/go-redis/v9"
)

var (
	dbRedis *redis.Client
)

func GetRedisOptions() Options {
	options := DefaultOptions
	options.Address = osenv.GetServerDbRedisAddress()
	return options
}

func Init() error {
	options := GetRedisOptions()
	dbRedis = redis.NewClient(&redis.Options{
		Addr:     options.Address,
		Password: options.Password,
		DB:       options.DB,
	})

	tctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := dbRedis.Ping(tctx).Err(); err != nil {
		return wrapErr(err)
	}

	repository.DBRedis = dbRedis
	return nil
}

func NewClient() (*redis.Client, error) {
	options := GetRedisOptions()

	client := redis.NewClient(&redis.Options{
		Addr:     options.Address,
		Password: options.Password,
		DB:       options.DB,
	})
	tctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(tctx).Err(); err != nil {
		return nil, wrapErr(err)
	}

	return client, nil
}

func Client() *redis.Client {
	return dbRedis
}

func Close() error {
	return dbRedis.Close()
}
