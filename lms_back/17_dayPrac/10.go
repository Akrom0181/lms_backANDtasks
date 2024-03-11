package main

import (
	"fmt"
)

func squared() {
	ch := make(chan int)
	sum := 0
	go func ()  {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for val := range ch {
		sum += val * val
	}
	fmt.Println("sum of squared num", sum)
}

func main(){
	squared()
}