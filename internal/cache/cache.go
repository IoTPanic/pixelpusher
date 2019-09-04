package cache

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// Constants for redis database indexes
const (
	OnlineDevices = 1
)

// NewPool allows for multiple proccesses to use the same connection
// https://medium.com/@gilcrest_65433/basic-redis-examples-with-go-a3348a12878e
func NewPool(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func InitalizePool(host string) {
	pool = NewPool(host)
	_, err := pool.Dial()
	if err != nil {
		panic("Redis connection failed")
	}
	log.Println("Redis pool connection created at", host)
}

// TerminatePool terminates the connection pool to the redis instance
func TerminatePool() {
	pool.Close()
}
