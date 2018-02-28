package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type Profile struct {
	Ops int `json:"Ops"`
}

func main() {
	http.HandleFunc("/", ops_service)
	http.ListenAndServe(":8082", nil)
}

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

func ops_service(w http.ResponseWriter, r *http.Request) {
	profile := Profile{benchmark()}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// CORS updates
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Request-Headers", "Access-Control-Allow-Origin")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(js)
}

func benchmark() int {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	mu.Lock()
	s, _ := redis.String(conn.Do("INFO", "stats"))
	stats_all := strings.Fields(s)
	ops_per := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
	ops_persec, _ := strconv.Atoi(ops_per)
	mu.Unlock()
	return ops_persec
}
