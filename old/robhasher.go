package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var mu sync.Mutex

// Global pool that handlers can grab a connection from
var pool = newPool()

// Pool configuration
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func loopcounter() int {
	conn := pool.Get()
	defer conn.Close()
	for i := 1; i <= 100; i++ {
		mu.Lock()
		for i := 0; i <= 200; i++ {
			conn.Do("INCR", "viewCount")
		}
		mu.Unlock()
		time.Sleep(1000 * time.Millisecond)
	}
	return 1
}

func main() {
	count := loopcounter()
	fmt.Println(count)
}
