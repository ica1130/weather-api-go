package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {

	limiter := rate.NewLimiter(2, 4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
