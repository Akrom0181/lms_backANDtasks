package main

import (
	"fmt"
)

func sumEven() {
	ch := make(chan int)
	sum := 0
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	go func() {
		for i := 1; i <= len(slice); i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		if num % 2 == 0 {
			sum += num
		}
	}

	fmt.Println("sum of evenNums:", sum)
}

func main() {
	sumEven()
}
