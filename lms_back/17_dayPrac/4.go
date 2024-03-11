package main

import (
	"fmt"
)

func isPrime(num int, primeChan chan<- bool) {
	if num <= 1 {
		primeChan <- false
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			primeChan <- false
			return
		}
	}
	primeChan <- true
}

func main() {
	var num int
	fmt.Print("Enter a number:")
	fmt.Scan(&num)
	primeChan := make(chan bool)

	go isPrime(num, primeChan)

	select {
	case isPrime := <-primeChan:
		if isPrime {
			fmt.Println("Prime number")
		} else {
			fmt.Println("Not a prime number")
		}
	}
}
