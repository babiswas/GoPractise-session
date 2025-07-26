package app

import (
	"fmt"
	"time"
)

func poolworker(n int, jobs <-chan int, results chan<- int) {
	for jb := range jobs {
		fmt.Println("Worker id:", n, "Started Job:", jb)
		time.Sleep(time.Second)
		fmt.Println("Worker id:", n, "Finished Job:", jb)
		results <- jb * 2
	}
}

func WorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < numJobs; i++ {
		go poolworker(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < numJobs; j++ {
		<-results
	}
}
