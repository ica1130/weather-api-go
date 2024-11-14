package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"weather-api-go.ilijakrilovic.net/internal/data"
	"weather-api-go.ilijakrilovic.net/internal/service"
)

func (app *application) weatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "city query parameter is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	cachedData, err := app.redis.Get(ctx, city).Result()

	if err == redis.Nil {

		weather, err := service.GetWeather(city)
		if err != nil {
			app.logger.Println(err)
			http.Error(w, fmt.Sprintf("error retrieving weather: %v", err), http.StatusInternalServerError)
			return
		}

		weatherJSON, err := json.Marshal(weather)

		if err != nil {
			app.logger.Println("error marshaling weather:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		err = app.redis.Set(ctx, city, weatherJSON, 10*time.Minute).Err()
		if err != nil {
			app.logger.Println("error saving data to redis:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		app.logger.Printf("new city added to redis: %s", city)
		app.writeJSON(w, http.StatusOK, weather, nil)

	} else if err != nil {

		app.logger.Println(err)
		http.Error(w, "error retrieving cached data", http.StatusInternalServerError)
		return

	} else {

		var weather data.Weather

		err = json.Unmarshal([]byte(cachedData), &weather)

		if err != nil {
			app.logger.Println("Error unmarshaling cached data:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		app.logger.Printf("city found: %s", city)
		app.writeJSON(w, http.StatusOK, weather, nil)
	}
}
