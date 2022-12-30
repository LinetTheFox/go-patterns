package main

import (
	"log"
	"os"
	"time"

	"example-runner/runner"
)

// timeout is the number of seconds the program has to finish
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work")

	// Create a new timer value for this to run
	r := runner.New(timeout)

	// Add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

// createTask returns an test task that sleeps for the specified
// number of seconds based on the id
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}