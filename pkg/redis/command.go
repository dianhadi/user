package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) Expire(ctx context.Context, key string, exp int64) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, exp)
	return err
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

func (r *Redis) HGetAll(ctx context.Context, key string) ([]interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}

	return values, err
}

func (r *Redis) HSet(ctx context.Context, key string, value interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", redis.Args{}.Add(key).AddFlat(value)...)
	return err
}

func (r *Redis) Set(ctx context.Context, key, value string) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

func (r *Redis) SetEx(ctx context.Context, key, value string, exp int64) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, exp, value)
	return err
}
