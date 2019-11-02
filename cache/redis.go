package cache

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(ip string, port int, passWord string, db int, dialTimeOut, readTimeout, writeTimeout,
poolTimeout time.Duration, poolSize int) *RedisCache {
	return &RedisCache{redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		Password:     passWord,
		DB:           db,
		DialTimeout:  dialTimeOut,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		PoolSize:     poolSize,
		PoolTimeout:  poolTimeout,
	})}
}

func (rc *RedisCache) Ping() bool {
	if err := rc.client.Ping().Err(); err != nil {
		return false
	}
	return true
}

func (rc *RedisCache) Set(key string, val interface{}, expiration time.Duration) (err error) {
	err = rc.client.Set(key, val, 0).Err()
	if expiration > 0 {
		rc.client.Expire(key, expiration)
	}
	return err
}

func (rc *RedisCache) GetString(key string) string {
	return rc.client.Get(key).String()
}

func (rc *RedisCache) GetInt64(key string) (val int64, err error) {
	return rc.client.Get(key).Int64()
}

func (rc *RedisCache) Get(key string, val interface{}) (ok bool, err error) {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		if err = rc.client.Get(key).Scan(val); err != nil {
			if err != redis.Nil {
				return false, err
			} else {
				return false, nil
			}
		} else {
			return true, nil
		}

	default:
		panic(errors.New("Need a pointer val "))
	}

}

func (rc *RedisCache) Exist(keys ...string) (int64, error) {
	return rc.client.Exists(keys...).Result()
}

func (rc *RedisCache) Expire(key string, t time.Duration) {
	rc.client.Expire(key, t)
}

func (rc *RedisCache) Delete(key ...string) (int64, error) {
	return rc.client.Del(key...).Result()
}

func (rc *RedisCache) HSet(key string, subKey string, val interface{}) (err error) {
	return rc.client.HSet(key, subKey, val).Err()
}

func (rc *RedisCache) HGet(key string, subKey string, val interface{}) (ok bool, err error) {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		if err = rc.client.HGet(key, subKey).Scan(val); err != nil {
			if err != redis.Nil {
				return false, err
			} else {
				return false, nil
			}
		} else {
			return true, nil
		}
	default:
		panic(errors.New("Need a pointer val "))
	}
}

func (rc *RedisCache) HGetAll(key string) (map[string]string, error) {
	return rc.client.HGetAll(key).Result()
}

func (rc *RedisCache) HMGet(key string, subKey ...string) ([]interface{}, error) {
	return rc.client.HMGet(key, subKey...).Result()
}

func (rc *RedisCache) HMSet(key string, fields map[string]interface{}) (err error) {
	return rc.client.HMSet(key, fields).Err()
}

func (rc *RedisCache) LRange(key string, start, stop int64) ([]string, error) {
	return rc.client.LRange(key, start, stop).Result()
}

func (rc *RedisCache) LPush(key string, data ...interface{}) (int64, error) {
	return rc.client.LPush(key, data...).Result()
}

func (rc *RedisCache) RPush(key string, data ...interface{}) (int64, error) {
	return rc.client.RPush(key, data...).Result()
}

func (rc *RedisCache) Close() error {
	return rc.client.Close()
}
