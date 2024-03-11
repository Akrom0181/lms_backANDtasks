package main

import (
    "fmt"
)

func reverseString(str string, result chan string) {
    reversed := ""
    for _, char := range str {
        reversed = string(char) + reversed
    }
    result <- reversed
}

func main() {

    result := make(chan string)
	fmt.Print("input txt: ")
    var str string
	fmt.Scan(&str) 

    go reverseString(str, result)

    reversed := <-result
    fmt.Println("Reversed:", reversed)
}
