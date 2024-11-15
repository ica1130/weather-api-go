package main

import (
	"net/http"

	"weather-api-go.ilijakrilovic.net/internal/service"
)

func (app *application) weatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "city query parameter is required", http.StatusBadRequest)
		return
	}

	weatherService := service.WeatherService{
		RedisClient: app.redis,
	}

	weather, err := weatherService.GetWeatherWithCache(r.Context(), city)
	if err != nil {
		app.logger.Println("error:", err)
		http.Error(w, "error retrieving weather", http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, weather, nil)
}
