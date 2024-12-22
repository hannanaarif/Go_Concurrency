package main

import ("fmt"; "sync")

var msg string
var wg sync.WaitGroup
var mu sync.Mutex

func updateMessage(s string,m *sync.Mutex){ 
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main(){

	msg="Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!",&mu)
	go updateMessage("Hello, cosmos!",&mu)
	wg.Wait()

	fmt.Println(msg)
}