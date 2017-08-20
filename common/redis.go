package common

import (
	"fmt"
	//	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

//func RedisSet(key, value string, ex time.Duration) {
//	err := client.Set(key, value, ex).Err()
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//func RedisGet(field, key string) string {
//	val, err := client.HGet(key, field).Result()
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//	return val
//}

//func RedisExistsKey(key string) *redis.IntCmd {
//	cmd := client.Exists(key)
//	return cmd
//}

//func RedisHexistsKey(field string, key string) *redis.BoolCmd {
//	cmd := client.HExists(key, field)
//	return cmd
//}
