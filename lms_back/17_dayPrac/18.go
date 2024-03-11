package main

import (
	"fmt"
	"strings"
)

func findLongestWord(text string, resultChan chan int) {
	words := strings.Fields(text)
	longest := 0

	for _, word := range words {
		wordLength := len(word)
		if wordLength > longest {
			longest = wordLength
		}
	}

	resultChan <- longest
}

func main() {

	text := "Create a program that uses goroutines and channels to find the length of the longest word in a given string concurrently."

	resultChan := make(chan int)

	go findLongestWord(text, resultChan)

	longestWordLength := <-resultChan

	fmt.Printf("The longest word in text: %d\n", longestWordLength)
}
