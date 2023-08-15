package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

var ErrNil = redis.ErrNil

func New(address string, port int, password string) (*Redis, error) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
	}

	return &Redis{pool: pool}, nil
}
