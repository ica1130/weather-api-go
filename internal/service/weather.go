package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"weather-api-go.ilijakrilovic.net/internal/data"
)

func GetWeather(city string) (*data.Weather, error) {

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
