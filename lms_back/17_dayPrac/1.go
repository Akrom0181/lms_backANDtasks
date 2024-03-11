package main

import (
	"fmt"
)

func printNumbers(ch chan int) {
	go func() {
		defer close(ch)

		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	printNumbers(ch)
}
