package main

import (
	"fmt"
	"net/http"

	"weather-api-go.ilijakrilovic.net/internal/data"
)

func (app *application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	var weather data.Weather

	err := app.readJSON(w, r, &weather)
	if err != nil {
		http.Error(w, "error occurred", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%+v\n", weather)
}
