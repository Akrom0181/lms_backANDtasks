package main

import (
	"fmt"
)

func findPrimesInRange(start, end int, primeChan chan<- int) {
	for num := start; num <= end; num++ {
		if num < 2 {
			continue
		}
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num % i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}
	}
	close(primeChan)
}

func main() {
	var start, end int
	fmt.Print("Enter the range start: ")
	fmt.Scan(&start)
	fmt.Print("Enter the range end: ")
	fmt.Scan(&end)

	primeChan := make(chan int)

	go findPrimesInRange(start, end, primeChan)

	fmt.Print("Prime numbers in the range: ")
	for prime := range primeChan {
		fmt.Print(prime, " ")
	}
	fmt.Println()
}
