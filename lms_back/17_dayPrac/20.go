package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(nums chan <- int, count int) {
	defer close(nums)
	for i := 0; i < count; i++ {
		nums <- rand.Intn(100)
	}
}

func consumer(nums <- chan int) {
	for val := range nums {
		fmt.Println("numbers: ", val)
	}
}

func main() {

	nums := make(chan int)

	go producer(nums, 10)
	go consumer(nums)

	time.Sleep(time.Second * 3)
	fmt.Println("timeSleep.....")
}