package app

import (
	"context"
	"fmt"
)

func generate_message(ctx context.Context, ch chan string, taskList []string) {
	for _, msg := range taskList {
		select {
		case ch <- msg:
			fmt.Println("Message sent.")
		case <-ctx.Done():
			fmt.Println("Task Cancelled.")
			return
		}
	}
}

func task_processor(ctx context.Context, taskList []string) chan string {
	ch := make(chan string)
	go generate_message(ctx, ch, taskList)
	return ch
}

func TaskExecutor(taskList []string) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := task_processor(ctx, taskList)
	for i := 0; i < 5000; i++ {
		fmt.Println(<-ch)
	}
	cancel()
}
