package main

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOSTNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	log.Print("Initialized redis client..")

	return client
}

func GetRedisClient() *redis.Client {
	return client
}
