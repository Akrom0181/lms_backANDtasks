package main

import (
	"fmt"
)

func factorial( ch chan int) {
	sum := 1
	var a int
	fmt.Println("kitiring")
	fmt.Scan(&a)
		go func() {
			defer close(ch)
			for i := 1; i <= a; i++ {
				ch <- i
			}
		}()
		for i := 1; i <= a; i++ {
			sum *= i
		}
		fmt.Println(sum)
}



func main() {
	new := make(chan int)
	factorial(new)
	
}
