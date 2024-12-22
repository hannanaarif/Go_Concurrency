package main

import (
	"fmt"
	"sync"
)

 type Income struct {
	source string
	amount int
}
var wg sync.WaitGroup

 func main() {
	var bankBalance int
	var balance sync.Mutex
	fmt.Printf("intial Bank balance: $%d.00\n", bankBalance)
    
	incomes:=[]Income{
		{source: "salary", amount: 1000},
		{source: "dividends", amount: 500},
		{source: "gift", amount: 200},
		{source: "lottery", amount: 1000},
		{source: "Investment", amount: 500},

	}
	wg.Add(len(incomes))

	for i,income:=range incomes{
		
		go func (i int,income Income){
			defer wg.Done()
			for week:=1;week<=52;week++{
				balance.Lock()
				temp:=bankBalance
				temp+=income.amount
				bankBalance=temp
				balance.Unlock()
				fmt.Printf("on week %d,you earned $%d.00 from %s\n",week,income.amount,income.source)
			}
	
		}(i,income)
	}
	wg.Wait()
	fmt.Printf("final bank balance: $%d.00\n",bankBalance)
}