package dbredis

import (
	"context"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/redis/go-redis/v9"
)

var (
	dbRedis      *redis.Client
	dbRedisSlave *redis.Client
)

func GetRedisOptions() Options {
	options := DefaultOptions
	options.Address = osenv.GetServerDbRedisAddress()
	if pw := osenv.GetServerDbRedisPassword(); len(pw) > 0 {
		options.Password = pw
	}
	options.DB = osenv.GetServerDbRedisDB()
	return options
}

func Init() error {
	if err := InitMaster(); err != nil {
		return err
	}
	if err := InitSlave(); err != nil {
		return err
	}
	return nil
}
func InitMaster() error {
	options := GetRedisOptions()
	dbRedis = redis.NewClient(&redis.Options{
		Addr:     options.Address,
		Password: options.Password,
		DB:       options.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := dbRedis.Ping(ctx).Err(); err != nil {
		return wrapErr(err)
	}

	//repository.DBRedis = dbRedis
	return nil
}
func InitSlave() error {
	options := DefaultOptions
	options.Address = osenv.GetServerDbRedisSlaveAddress()
	if pw := osenv.GetServerDbRedisSlavePassword(); len(pw) > 0 {
		options.Password = pw
	}
	options.DB = osenv.GetServerDbRedisSlaveDB()
	dbRedisSlave = redis.NewClient(&redis.Options{
		Addr:     options.Address,
		Password: options.Password,
		DB:       options.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := dbRedisSlave.Ping(ctx).Err(); err != nil {
		return wrapErr(err)
	}

	//repository.DBRedis = dbRedis
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
func ClientSlave() *redis.Client {
	return dbRedisSlave
}

func Close() error {
	return dbRedis.Close()
}
