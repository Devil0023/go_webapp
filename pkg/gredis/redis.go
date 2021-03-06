package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_webapp/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

//Setup 注册Redis
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

//Set set命令
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

//Get get命令
func Get(key string) ([]byte, error) {

	conn := RedisConn.Get()

	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", key))

	if err != nil {
		return nil, err
	}

	return result, nil
}

//Exists exists命令
func Exists(key string) bool {

	conn := RedisConn.Get()

	defer conn.Close()

	result, err := redis.Bool(conn.Do("EXISTS", key))

	if err != nil {
		return false
	}

	return result
}

//Delete delete命令
func Delete(key string) (bool, error) {

	conn := RedisConn.Get()

	defer conn.Close()

	return redis.Bool(conn.Do("DELETE", key))
}

//Ttl ttl命令
func Ttl(key string) (int64, error) {
	conn := RedisConn.Get()

	defer conn.Close()

	return redis.Int64(conn.Do("TTL", key))
}

//HGetAll hgetall
func HGetAll(key string) (map[string]interface{}, error) {

	conn := RedisConn.Get()

	data := make(map[string]interface{})

	defer conn.Close()

	slices, err := redis.ByteSlices(conn.Do("HGETALL", key))

	if err != nil {
		return nil, err
	}

	for i, v := range slices {

		if i%2 == 0 {
			data[string(v)] = string(slices[i+1])
		}

	}

	return data, nil

}
