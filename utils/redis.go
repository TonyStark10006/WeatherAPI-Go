package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func getAValue(key string) (value string, err error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		panic(err)
	}

	return val, err
}

func setAValue(key string, value interface{}) (err error) {
	err = client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	} else {
		err = nil
	}

	return err
}
