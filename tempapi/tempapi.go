package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
)

type Hashes struct {
	hash1 int `json:"hashes"`
}

type CoinCount struct {
	coins1 int `json:"coins"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hashes", hashesPS)
	router.HandleFunc("/coins", coinReturn)
	log.Fatal(http.ListenAndServe(":8082", router))
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

func coinReturn(w http.ResponseWriter, r *http.Request) {
	coins := CoinCount{totCoins()}

	js1, err := json.Marshal(coins)
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
	w.Write(js1)
}

func hashesPS(w http.ResponseWriter, r *http.Request) {
	hashes := Hashes{benchmark()}

	js, err := json.Marshal(hashes)
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

func totCoins() int {
	conn := pool.Get()
	defer conn.Close()
	count, _ := redis.Int(conn.Do("GET", "viewCount"))
	return count
}

func benchmark() int {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	s, _ := redis.String(conn.Do("INFO", "stats"))
	stats_all := strings.Fields(s)
	ops_per := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
	ops_persec, _ := strconv.Atoi(ops_per)
	return ops_persec
}
