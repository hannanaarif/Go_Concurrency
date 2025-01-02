package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, ch chan int) {
	for task := range ch {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	taskChannels := make(chan chan int) // Channel of channels

	// Start a manager goroutine to create task channels
	go func() {
		for i := 1; i <= 3; i++ {
			taskChan := make(chan int)   // Create a channel for each worker
			taskChannels <- taskChan     // Send taskChan to taskChannels
			go worker(i, taskChan)       // Start a worker for this taskChan
		}
		close(taskChannels) // Close taskChannels when done
	}()

	// Main goroutine sends tasks to workers via taskChannels
	for taskChan := range taskChannels {
		for task := 1; task <= 3; task++ {
			taskChan <- task // Send tasks to the worker's taskChan
		}
		close(taskChan) // Close the individual taskChan after sending tasks
	}

	fmt.Println("All tasks sent!")
}
