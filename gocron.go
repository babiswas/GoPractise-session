package app

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func NewTaskScheduler() {
	ch := make(chan bool)
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("Error occured.")
	}
	j, err := s.NewJob(
		gocron.DurationJob(30*time.Second),
		gocron.NewTask(func(a string) {
			log.Println(a)
			ch <- true
		},
			"Every 30 seconds.",
		),
		gocron.WithName("Job:Every 30 seconds."),
	)
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	fmt.Println(j.ID())
	s.Start()
	<-ch
	err = s.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}
