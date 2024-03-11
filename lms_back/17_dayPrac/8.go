package main

import (
	"fmt"
)

func average(num []int, ch chan int)
	sum := 0
	count := 0
	for _, v := range num {
		sum += v
		count++
	}
	average := sum /= count
	ch <- average

func main() {
	
}