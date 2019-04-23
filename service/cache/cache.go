package cache

import (
	//"log"
	db "github.com/dangyanglim/go_cnode/database"
	"github.com/garyburd/redigo/redis"
)

func Get(key string) (interface{}, error) {
	temp, err := redis.Bytes(db.Redis.Do("GET", key))
	return temp, err
}
func Set(key string, data interface{}) error {
	_, err := redis.String(db.Redis.Do("SET", key, data))

	return err
}
func SetEx(key string, data interface{}) error {
	_, err := redis.String(db.Redis.Do("SET", key, data, "EX", 60))

	return err
}
