package main

import (
	"fmt"

	"github.com/prprprus/scheduler"
)

/*
 Purpose: Scheduler
	 - job scheduling packag
	 - Installation command
		 go get github.com/prprprus/scheduler
	 - Features
		- Delay execution, accurate to a second
		- Periodic execution, accurate to a second, like the cron style but more flexible
		- Cancel job
		- Failure retry
*/

func main() {
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		panic(err) // just example
	}

	// delay with 1 second, job function with arguments
	s.Delay().Second(1).Do(task1, "prprprus", 23)

	// delay with 1 minute, job function without arguments
	s.Delay().Minute(1).Do(task2)

	// delay with 1 hour
	s.Delay().Hour(1).Do(task2)

	// special: execute immediately
	s.Delay().Do(task2)

	// cancel job
	jobID := s.Delay().Day(1).Do(task2)
	err = s.CancelJob(jobID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("cancel delay job success")
	}
}
