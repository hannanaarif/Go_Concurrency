package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var NumberofPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan pizzaOrder
	quit chan chan error
}

type pizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch

}
func makePizza(pizzaNumber int) *pizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberofPizzas {

		delay := rand.Intn(5) + 1
		fmt.Printf("Received the order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		fmt.Printf("Making pizza #%d.It will be ready in %d seconds\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingredients for pizza #%d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("The cook quit while making pizza #%d\n", pizzaNumber)
		} else {
			msg = fmt.Sprintf("Pizza #%d is ready!\n", pizzaNumber)
			success = true
		}
		p := pizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &pizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func Pizzeria(pizzaMaker *Producer) {

	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			//we tried to make a pizza(we sent something to data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}

		}

	}
}

func main() {
	//seed the random number generator
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan pizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in a goroutine/bacground thread
	go Pizzeria(pizzaJob)

	//create and run consumer

	for i := range pizzaJob.data {

		if i.pizzaNumber <= NumberofPizzas {

			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order #%d has been cancelled\n", i.pizzaNumber)
			}
		}else{
			color.Cyan("The Pizzeria is closed for the day!")
			err:=pizzaJob.Close()
			if err!=nil{
				color.Red("Error closing the pizzeria",err)
			}
		}
       //print out the ending message
	}
}
