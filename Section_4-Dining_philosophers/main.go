package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eattime = 1 * time.Second
var thinktime = 0 * time.Second
var sleep = 1 * time.Second

func main() {
	//The Dining Philosophers Problem
	//five philosphers from 0 to 4
	//five forks from 0 to 4
	//there only difficulty -besides those of philosophy is that dishes served is very difficult kind of spaghetti which has to eaten
	//with two forks.There are two forks next to each plate, and, in order to eat, a philosopher must have both forks.

	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	fmt.Println("The table is empty")

	dine()

	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(5)

	seated := &sync.WaitGroup{}
	seated.Add(5)

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		fmt.Printf("%s has sat down\n", philosophers[i].name)
		go diningProblem(philosophers[i], forks, wg, seated)
	}
	wg.Wait()
}

func diningProblem(philosopher Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()
	seated.Done()
	seated.Wait()

	for i := 0; i < hunger; i++ {

		if philosopher.leftFork > philosopher.rightFork {

			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t %s has picked up right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t %s has picked up left fork\n", philosopher.name)

		} else {

			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t %s has picked up left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t %s has picked up right fork\n", philosopher.name)
		}

		fmt.Printf("\t %s has both the fork and is eating\n", philosopher.name)
		time.Sleep(eattime)

		fmt.Printf("\t %s is thinking\n", philosopher.name)
		time.Sleep(thinktime)

		fmt.Printf("\t %s has finished eating and is putting down the fork\n", philosopher.name)
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
	}
	fmt.Printf("%s is satisfied \n", philosopher.name)
	fmt.Printf("%s has left the table\n", philosopher.name)
}
