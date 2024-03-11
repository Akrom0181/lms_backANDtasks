package main

import (
	"fmt"
)

func multiply(slice []int, ch chan []int) {
    var newSlice []int
    for _, v := range slice {
        newSlice = append(newSlice, v*2)
    }
    ch <- newSlice
}

func main(){
	result := make(chan []int)

    numbers := []int{4, 23, 9, 12, 34, 8, 67, 55, 3, 1} 

	go multiply(numbers, result)
	mu2x := <- result
	fmt.Println(mu2x)
}