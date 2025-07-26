package app

import (
	"fmt"
	"sync"
	"time"
)

// Worker function.
func worker(id int) {
	fmt.Println("Starting workers:", id)
	time.Sleep(time.Second)
	fmt.Println("Work completed:", id)
}

// Execute workers.
func ExecuteWorkers() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
