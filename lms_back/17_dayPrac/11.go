package main

import (
	"fmt"
)

func squaring1000() {
	ch := make(chan int)
	sum := 0

	go func() {
		for i := 1 ; i <= 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		sum += v * v
	}
	fmt.Println("Yig'indisi: ", sum)
}

func main(){
	squaring1000()
}