package main

import "github.com/garyburd/redigo/redis"
import "fmt"
import "strings"
import "time"

// increment counter by one and return total counter value
func benchmark() string {
	// Grab a connection and make sure to close it with defer
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}

	defer c.Close()
	s, _ := redis.String(c.Do("INFO", "stats"))
	stats_all := strings.Fields(s)
	ops_persec := strings.TrimPrefix(stats_all[4], "instantaneous_ops_per_sec:")
	return ops_persec
}

func main() {

	for i := 1; i <= 100; i++ {
		ops := benchmark()
		fmt.Println(ops)
		time.Sleep(3 * time.Second)
	}
	fmt.Println("End of report")
}
