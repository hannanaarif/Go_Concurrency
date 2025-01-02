package main

import (
	"fmt"
	"time"
)

func worker1(id int, taskChan chan int, nextChan chan int) {
	for task := range taskChan {
		fmt.Printf("Worker 1 (ID: %d) processing task: %d\n", id, task)
		// Simulate some processing
		time.Sleep(time.Millisecond * 500)
		// Send task to the next stage (worker 2)
		nextChan <- task + 1
	}
	close(nextChan) // Close the next channel when done
}

func worker2(id int, taskChan chan int, nextChan chan int) {
	for task := range taskChan {
		fmt.Printf("Worker 2 (ID: %d) processing task: %d\n", id, task)
		// Simulate more processing
		time.Sleep(time.Millisecond * 500)
		// Send task to the next stage (worker 3)
		nextChan <- task * 2
	}
	close(nextChan) // Close the next channel when done
}

func worker3(id int, taskChan chan int) {
	for task := range taskChan {
		fmt.Printf("Worker 3 (ID: %d) final result: %d\n", id, task)
	}
}

func main() {
	// Create the outermost channel of channels (chan chan int)
	taskChannels := make(chan chan int, 3)

	// Create channels for each stage and start workers
	for i := 1; i <= 3; i++ {
		// Create the second level of channels
		taskChan := make(chan int)  // Stage 1's channel
		nextChan1 := make(chan int) // Stage 2's channel
		nextChan2 := make(chan int) // Stage 3's channel

		go worker1(i, taskChan, nextChan1)  // Start Stage 1 (Worker 1)
		go worker2(i, nextChan1, nextChan2) // Start Stage 2 (Worker 2)
		go worker3(i, nextChan2)            // Start Stage 3 (Worker 3)

		// Send the task channel to the outer channel (taskChannels)
		taskChannels <- taskChan
	}

	// Send tasks through the stages
	for taskChan := range taskChannels {
		for task := 1; task <= 1; task++ {
			taskChan <- task // Send tasks to worker 1's task channel
		}
		close(taskChan) // Close taskChan after sending tasks
	}

	// Wait for all workers to finish processing
	time.Sleep(time.Second * 3)
	fmt.Println("Pipeline processing completed!")
}
