package db

import "github.com/TangSengDaoDao/TangSengDaoDaoServer/pkg/redis"

func NewRedis(addr string) *redis.Conn {
	return redis.New(addr)
}
