package config

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	connStr := fmt.Sprintf("%s:%s", REDISHost, REDISPort)
	var addr = flag.String("Server", connStr, "Redis server address")
	fmt.Println("Successful connected to Redis:", string(*addr))

	rdb := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
