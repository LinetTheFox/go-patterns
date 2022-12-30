package main

import (
	"log"
	"sync"
	"time"

	"example-work/work"
)

// names provides a set of names to display
var names = []string{
	"Steve",
	"Bob",
	"Mary",
	"Therese",
	"Jason",
}

// namePrinter provides special support for printing names
type namePrinter struct {
	name string
}

// Task implements the Worker interface
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// Create a work pool with 2 goroutines
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names
		for _, name := range names {
			// Create a namePriter and provide
			// the specific name
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When
				// RunTask returns we know it is being
				// handled
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shut down the work pool and wait for all existing work
	// to be completed
	p.Shutdown()
}
