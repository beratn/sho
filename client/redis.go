package client

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
)

var client *redis.Client

func InitRedis() *redis.Client {
	log.Print("Initializing redis client..")

	if client != nil {
		return client
	}

	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	log.Print("Initialized redis client..")

	return client
}

func GetRedisClient() *redis.Client {
	return client
}
