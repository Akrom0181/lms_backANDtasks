package main

import (
	"fmt"
	"strings"
)

func countWords(text string, resultChan chan<- int) {
	words := strings.Fields(text)
	wordCount := len(words)
	resultChan <- wordCount
}

func main() {
	text := "this exercise is fourteenth one"

	resultChan := make(chan int)

	go countWords(text, resultChan)

	wordCount := <-resultChan

	fmt.Println("count of words in text:", wordCount)
}
