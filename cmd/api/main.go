package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	redis "github.com/redis/go-redis/v9"
	redisConfig "weather-api-go.ilijakrilovic.net/internal/redis"
)

const version = "1.0.0."

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	redis  *redis.Client
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	redisClient, err := redisConfig.InitializeRedis()

	if err != nil {
		log.Fatalf("error establishing a redis connection: %v", err)
		os.Exit(1)
	}

	defer redisClient.Close()

	app := &application{
		config: cfg,
		logger: logger,
		redis:  redisClient,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
