package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	// fmt.Println("hello concurrency")
	// go printSomething("this is my first thing to be printed")
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d:%s", i, x), &wg)
	}
	wg.Wait()

	// time.Sleep(1*time.Second)
	wg.Add(1)
	printSomething("this is my second thing to be printed", &wg)
}
