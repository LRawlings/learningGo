package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDurration := measureResponse(a)
	bDurration := measureResponse(b)

	if aDurration < bDurration {
		return a
	}

	return b
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
