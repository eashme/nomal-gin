package cache

import "time"

type ICache interface {
	Set(key string, val interface{}, expiration time.Duration) (err error)
	GetString(key string) string
	GetInt64(key string) (val int64, err error)
	Get(key string, val interface{}) (ok bool, err error)
	Exist(keys ...string) (int64, error)
	Expire(key string, t time.Duration)
	Delete(key ...string) (int64, error)
	HSet(key string, subKey string, val interface{}) (err error)
	HGet(key string, subKey string, val interface{}) (ok bool, err error)
	HGetAll(key string) (map[string]string, error)
	HMGet(key string, subKey ...string) ([]interface{}, error)
	HMSet(key string, fields map[string]interface{}) (err error)
	LPush(key string, data ...interface{}) (int64, error)
	RPush(key string, data ...interface{}) (int64, error)
	LRange(key string, start, stop int64) ([]string, error)
	Close() error
}
