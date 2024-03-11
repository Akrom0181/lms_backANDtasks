package main

import (
    "fmt"
)

func sum(result chan int) {
    go func() {
        defer close(result)
        sum := 0
        for i := 1; i <= 100; i++ {
            sum += i
        }
        result <- sum
    }()
}

func main() {
    resultChan := make(chan int)
    sum(resultChan)

    result := <-resultChan
    fmt.Println(result)
}
