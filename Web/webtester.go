package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type Page struct {
	Body  string
	Title string
	Name  string
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

//func ignoreIcon(res http.ResponseWriter, req *http.Request) {
//	http.ServeFile(res, req, "favicon.ico")
// }

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("index.html") //create a new template with some name
	tmpl, _ = tmpl.ParseFiles("index.html")

	bench := benchmark()
	bodString := fmt.Sprintf(" %v RobCoins have been mined ", bench)
	p := Page{Body: bodString, Title: `RobCoin Miner`, Name: "Robert"}

	if err := tmpl.Execute(w, p); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

// increment counter by one and return total counter value
func benchmark() string {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	mu.Lock()
	s, _ := redis.String(conn.Do("INFO", "stats"))
	stats_all := strings.Fields(s)
	ops_persec := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
	hname, _ := os.Hostname()
	mu.Unlock()
	bench := hname + " " + ops_persec
	hname == ""
	return bench
}

func main() {
	//http.HandleFunc("/favicon.ico", ignoreIcon)
	http.HandleFunc("/", serveTemplate)
	log.Println("Server Listening...")
	http.ListenAndServe(":8082", nil)
}
