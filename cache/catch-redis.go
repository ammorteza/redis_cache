package cache

import (
	"github.com/gomodule/redigo/redis"
)

type redisCache struct {
	conn 			redis.Conn
	db 				int
	exp 			int
}

func NewRedisCache(db int, exp int) ICache{
	return &redisCache{
		conn: pool().Get(),
		db: db,
		exp: exp,
	}
}

func (this *redisCache)Set(key string, val string) error{
	_, err := this.conn.Do("SET" , key, val, "EX", this.exp)
	if err != nil{
		return err
	}
	return nil
}

func (this *redisCache)Get(key string) (string, error){
	return redis.String(this.conn.Do("GET" , key))
}