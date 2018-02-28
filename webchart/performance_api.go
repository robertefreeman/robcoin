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
	pvalue  string
	pnumber string
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

// increment counter by one and return total counter value
func benchmark() int {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	mu.Lock()
	s, _ := redis.String(conn.Do("INFO", "stats"))
	stats_all := strings.Fields(s)
	ops_per := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
	ops_persec, _ := strconv.Atoi(ops_per)
	// hname, _ := os.Hostname()
	mu.Unlock()
	// bench := hname + " " + ops_persec
	// hname == ""
	return ops_persec
}

//Javascript array creation function

func outputOps(w http.ResponseWriter, r *http.Request) {
	// message := r.URL.Path
	// message = strings.TrimPrefix(message, "/")
	// message = "Hello " + message
	//profile := Profile{"plot0", benchmark()}
	mapD := map[string]int{"plot0": benchmark()}
	mapB, _ := json.Marshal(mapD)
	// message := "[\"plot0\":" + benchmark() + "]"
	//opsvalue, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write(mapB)
}

func main() {
	http.HandleFunc("/", outputOps)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
