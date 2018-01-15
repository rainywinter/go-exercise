package cmd

// redis 事务加watch 测试

import (
	"math/rand"
	"rw/go-exercise/redis/db"
	"time"

	"github.com/golang/glog"

	"github.com/go-redis/redis"
)

var (
	watchKey = "watch_transaction"
)

func init() {
	client := db.GetRedis()

	err := client.Watch(func(tx *redis.Tx) error {
		_, err := tx.Get(watchKey).Int64()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set("weather", "sunny", 10*time.Second)
			// 同一个链接内改变watch的key不影响逻辑
			pipe.Set(watchKey, rand.Intn(100), 10*time.Second)
			return nil
		})
		return err
	}, watchKey)
	if err != nil {
		glog.Errorf("key changed: %v", err)
	}
}
