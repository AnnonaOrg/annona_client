package db_data

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/AnnonaOrg/annona_client/internal/dbredis"
	"github.com/redis/go-redis/v9"
)

// NilErr represents redis nil
var NilErr = errors.New("redis.Nil")

func wrapErr(err error) error {
	if err == redis.Nil {
		return NilErr
		// return nil
	}
	return err
}
func RDBClient() *redis.Client {
	return dbredis.Client()
	// return db.RDB
}
func RDBClientSlave() *redis.Client {
	return dbredis.ClientSlave()
	// return db.RDB
}

// addToSetWithExpiration 向集合中添加元素
func AddToSet(key string, member interface{}) error {
	client := RDBClient()
	// 使用SAdd方法添加元素到集合
	_, err := client.SAdd(context.Background(), key, member).Result()
	return wrapErr(err)
}

// addMultipleToSet 批量添加元素到集合
func AddMultipleToSet(key string, members ...interface{}) error {
	client := RDBClient()
	// 使用SAdd方法添加元素到集合
	_, err := client.SAdd(context.Background(), key, members...).Result()
	return wrapErr(err)
}

// addToSetWithExpiration 向集合中添加元素并设置过期时间
func AddToSetWithExpiration(key string, member interface{}, expiration time.Duration) error {
	client := RDBClient()
	// 使用事务保证原子性
	_, err := client.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		// 向集合中添加元素
		pipe.SAdd(context.Background(), key, member)
		// 设置过期时间
		pipe.Expire(context.Background(), key, expiration)
		return nil
	})
	return wrapErr(err)
}

// addToSetWithExpiration 批量添加元素到集合并设置过期时间
func AddMultipleToSetWithExpiration(key string, expiration time.Duration, members ...interface{}) error {
	client := RDBClient()
	// 使用事务保证原子性
	_, err := client.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		// 向集合中批量添加元素
		pipe.SAdd(context.Background(), key, members...)
		// 设置过期时间
		pipe.Expire(context.Background(), key, expiration)
		return nil
	})
	return wrapErr(err)
}

// removeFromSet 从集合中删除元素
func RemoveFromSet(key string, member interface{}) error {
	client := RDBClient()
	// 使用事务保证原子性
	_, err := client.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		// 从集合中删除元素
		pipe.SRem(context.Background(), key, member)
		return nil
	})
	return wrapErr(err)
}

// deleteAllSetElements 删除集合中所有元素
func RemoveAllFromSet(key string) error {
	client := RDBClient()
	// 使用Del方法删除集合键
	return client.Del(context.Background(), key).Err()
}

// getSetMembers 获取集合中的所有元素
func GetSetMembers(key string) ([]string, error) {
	client := RDBClientSlave()
	// 使用SMembers方法获取集合中的所有元素
	result, err := client.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, wrapErr(err)
	}
	return result, nil
}

// isMemberOfSet 检查元素是否属于集合
func IsMemberOfSet(key string, member interface{}) (bool, error) {
	client := RDBClientSlave()
	// 使用SIsMember方法检查元素是否属于集合
	result, err := client.SIsMember(context.Background(), key, member).Result()
	if err != nil {
		return false, wrapErr(err)
	}
	return result, nil
}

// setStructWithExpiration 添加带有过期时间的结构体到键
func AddKeyValueWithExpiration(key string, value interface{}, expiration time.Duration) error {
	client := RDBClient()
	// 序列化结构体为JSON字符串
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return wrapErr(err)
	}

	// 使用Set方法添加带有过期时间的字符串到键
	err = client.Set(context.Background(), key, jsonValue, expiration).Err()
	return wrapErr(err)
}

// setKeyValueWithoutExpiration 添加键值对到Redis（不携带过期时间）
func AddKeyValue(key, value string) error {
	client := RDBClient()
	// 序列化结构体为JSON字符串
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return wrapErr(err)
	}

	// 使用Set方法添加字符串到键
	err = client.Set(context.Background(), key, jsonValue, 0).Err()
	return wrapErr(err)
}

// getStruct 获取键的结构体值
func GetKeyValue(key string, value interface{}) error {
	client := RDBClientSlave()
	// 使用Get方法获取键的字符串值
	jsonValue, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return wrapErr(err)
	}

	// 反序列化JSON字符串为结构体
	return json.Unmarshal([]byte(jsonValue), value)
}

func DelKeyValue(key string) (int64, error) {
	client := RDBClient()
	v, err := client.Del(context.Background(), key).Result()

	return v, wrapErr(err)
}
