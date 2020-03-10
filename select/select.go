package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDurration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDurration := time.Since(startB)

	if aDurration < bDurration {
		return a
	}

	return b
}
