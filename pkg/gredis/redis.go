package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_webapp/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {

	RedisConn = &redis.Pool{

		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,

		Dial: func() (conn redis.Conn, err error) {

			conn, err = redis.Dial("tcp", setting.RedisSetting.Host)

			if err != nil {
				return nil, err
			}

			if setting.RedisSetting.Password != "" {
				if _, err = conn.Do("AUTH", setting.RedisSetting.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}

			if setting.RedisSetting.Database != 0 {
				if _, err = conn.Do("SELECT", setting.RedisSetting.Database); err != nil {
					conn.Close()
					return nil, err
				}
			}

			return conn, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (string, error) {

	var result string

	conn := RedisConn.Get()

	defer conn.Close()

	value, err := json.Marshal(data)

	if err != nil {
		return result, err
	}

	if time != 0 {
		result, err = redis.String(conn.Do("SET", key, value, "EX", time))
	} else {
		result, err = redis.String(conn.Do("SET", key, value))
	}

	return result, err
}

func Get(key string) ([]byte, error) {

	conn := RedisConn.Get()

	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", key))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Exists(key string) bool {

	conn := RedisConn.Get()

	defer conn.Close()

	result, err := redis.Bool(conn.Do("EXISTS", key))

	if err != nil {
		return false
	}

	return result
}

func Delete(key string) (bool, error) {

	conn := RedisConn.Get()

	defer conn.Close()

	return redis.Bool(conn.Do("DELETE", key))
}
