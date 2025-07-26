package app

import (
	"fmt"
	"time"
)

// Request Limiter.
func RequestLimiter(n int) {
	request := make(chan int, n)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)
	limiter := time.NewTicker(400 * time.Millisecond)
	defer limiter.Stop()

	for req := range request {
		<-limiter.C
		fmt.Println("Request:", req, time.Now())
	}
}

// Bursty Limiter.
func BurstyLimiter(n int) {
	burstyLimiter := make(chan time.Time, 3)
	for range 3 {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	burstyRequest := make(chan int, n+1)
	fmt.Println("Starting bursty request.")
	for i := 0; i <= n; i++ {
		burstyRequest <- i
	}
	fmt.Println("Closing bursty request:")
	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("Request:", req, time.Now())
	}

}
