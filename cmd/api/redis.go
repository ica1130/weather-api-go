package main

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisCtx = context.Background()

func initializeRedis() (*redis.Client, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *redis.Client

	for {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		_, err := rdb.Ping(redisCtx).Result()
		if err != nil {
			log.Println("redis is not ready yet")
			counts++
		} else {
			log.Println("succesfully connected to redis")
			connection = rdb
			break
		}

		if counts > 5 {
			log.Fatalf("could not connect to redis: %v", err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off..")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
