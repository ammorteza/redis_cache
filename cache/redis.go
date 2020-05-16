package cache

import "github.com/gomodule/redigo/redis"

func pool() *redis.Pool{
	return &redis.Pool{
		MaxActive: 12000,
		MaxIdle: 80,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}