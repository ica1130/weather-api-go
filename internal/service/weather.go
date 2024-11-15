package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"weather-api-go.ilijakrilovic.net/internal/data"
)

type WeatherService struct {
	RedisClient *redis.Client
}

func (w WeatherService) GetWeatherWithCache(ctx context.Context, city string) (*data.Weather, error) {

	var weather *data.Weather
	expiration := 10 * time.Minute

	cachedData, err := w.RedisClient.Get(ctx, city).Result()

	if err == redis.Nil {
		weather, err := getWeatherFromAPI(city)
		if err != nil {
			return weather, err
		}

		weatherJSON, err := json.Marshal(weather)
		if err != nil {
			return weather, err
		}

		err = w.RedisClient.Set(ctx, city, weatherJSON, expiration).Err()
		return weather, err
	} else if err != nil {
		return weather, err
	} else {
		err = json.Unmarshal([]byte(cachedData), &weather)
		return weather, err
	}

}

func getWeatherFromAPI(city string) (*data.Weather, error) {

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("weather API key is not set")
	}

	apiURL := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", city, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status: %s", resp.Status)
	}

	var weather data.Weather

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, fmt.Errorf("failed to decode weather API response: %w", err)
	}

	return &weather, nil

}
