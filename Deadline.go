package app

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func task_executor(done chan bool, long_running_task func()) {
	num := rand.Intn(5)
	if num%2 == 0 {
		fmt.Println("Even Number:")
		done <- true
	} else {
		long_running_task()
	}
}

func long_running_task() {
	for {
		fmt.Println("Executing task.")
		time.Sleep(2 * time.Millisecond)
		fmt.Println("Completing task.")
	}
}

func Start_Execution() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Millisecond))
	defer cancel()
	done := make(chan bool)
	go task_executor(done, long_running_task)
	select {
	case <-done:
		fmt.Println("Task Sucessful.")
	case <-ctx.Done():
		fmt.Println("Task Cancelled.")
	}
}
