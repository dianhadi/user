package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *Redis) ScanStruct(source []interface{}, target interface{}) error {
	return redis.ScanStruct(source, target)
}
