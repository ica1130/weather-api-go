package main

import (
	"fmt"
	"net/http"

	"weather-api-go.ilijakrilovic.net/internal/service"
)

func (app *application) weatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "city query parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := service.GetWeather(city)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, fmt.Sprintf("error retrieving weather: %v", err), http.StatusInternalServerError)
		return
	}

	err = app.writeJSON(w, http.StatusOK, weather, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server ecnountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
