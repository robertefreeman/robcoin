package main

import (
	"fmt"
	"math/rand"
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
			return redis.Dial("tcp", "redis:6379")
		},
	}
}

func coinhash() int {
	conn := pool.Get()
	defer conn.Close()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		hash_rand := r1.Intn(100)
		if hash_rand == 18 {
			mu.Lock()
			conn.Do("INCR", "viewCount")
			mu.Unlock()
			coinfind()
		}
		time.Sleep(1 * time.Millisecond)
	}
	return 1
}

func coinfind() {
	conn := pool.Get()
	defer conn.Close()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	hash_rand := r1.Intn(100)
	if hash_rand == 3 {
		mu.Lock()
		conn.Do("INCR", "coins")
		mu.Unlock()
		count, _ := redis.Int(conn.Do("GET", "coins"))
		fmt.Println(count)
	}
}

func main() {
	count := coinhash()
	fmt.Println(count)
}
