package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{4, 9, 16, 25, 36}

	resultChan := make(chan float64, len(numbers))

	for _, num := range numbers {
		go func(num int) {
			squareRoot := math.Sqrt(float64(num))
			resultChan <- squareRoot
		}(num)
	}

	for range numbers {
		squareRoot := <-resultChan
		fmt.Printf("Square root: %.2f\n", squareRoot)
	}
}
