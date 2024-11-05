package main

import (
	"fmt"
	"net/http"
)

func (app *application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get weather data")
}
