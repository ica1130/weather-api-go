package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var redisCtx = context.Background()

func initializeRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(redisCtx).Result(); err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	} else {
		log.Println("succesfully connected to redis")
	}

	return rdb
}
