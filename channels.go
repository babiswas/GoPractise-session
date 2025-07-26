package app

import "fmt"

func JobCreator() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("REceived all jobs.")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent jobs:", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs:")
	<-done
	_, ok := <-jobs
	fmt.Println("Recieved more jobs:", ok)
}
