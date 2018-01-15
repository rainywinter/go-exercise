package db

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/go-redis/redis"
)

var (
	host   = "127.0.0.1"
	port   = 6379
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%d", host, port),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		glog.Fatalf("dial redis connection failed,pong=%v,err=%v", pong, err)
	} else {
		glog.Info("conn redis success.")
	}
}

// GetRedis return redis client
func GetRedis() *redis.Client {
	return client
}
