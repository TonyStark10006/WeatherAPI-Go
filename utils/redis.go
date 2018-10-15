package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	if client == nil {
		loadErr := godotenv.Load()
		host := os.Getenv("REDIS_HOST")
		pwd := os.Getenv("REDIS_PASSWORD")
		port := os.Getenv("REDIS_PORT")
		client = redis.NewClient(&redis.Options{
			Addr:     host + ":" + string(port),
			Password: pwd, // no password set
			DB:       0,   // use default DB
		})

		pong, err := client.Ping().Result()
		if err != nil {
			fmt.Println("尝试ping redis服务出错:" + err.Error())
			log.Fatal("尝试ping redis服务出错:" + err.Error())
		}
		fmt.Println(pong, err, loadErr)
	}

}

func GetAValue(key string) (value string, err error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		panic(err)
	}

	return val, err
}

func SetAValue(key string, value interface{}) (err error) {
	err = client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	} else {
		err = nil
	}

	return err
}
