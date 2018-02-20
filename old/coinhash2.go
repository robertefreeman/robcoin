package main

import (
	"fmt"
	"os"
	"strings"
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

func main() {
	conn := pool.Get()
	defer conn.Close()
	for i := 0; i <= 10; i++ {
		mu.Lock()
		s, _ := redis.String(conn.Do("INFO", "stats"))
		stats_all := strings.Fields(s)
		ops_persec := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
		mu.Unlock()
		fmt.Println(ops_persec)
		hname, _ := os.Hostname()
		fmt.Println(hname)
		time.Sleep(3 * time.Second)
	}
}
