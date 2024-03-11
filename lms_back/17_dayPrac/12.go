package main

import (
	"fmt"
	"strings"
)

func VowelsAndConsonants(text string, vowelCountCh, consonantCountCh chan<- int) {
	vowels := "aeiouAEIOU"
	vowelCount := 0
	consonantCount := 0

	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			consonantCount++
		}
	}

	vowelCountCh <- vowelCount
	consonantCountCh <- consonantCount
}

func main() {
	text := "Assalom aleykum"

	vowelCountCh := make(chan int)
	consonantCountCh := make(chan int)

	go VowelsAndConsonants(text, vowelCountCh, consonantCountCh)

	vowels := <-vowelCountCh
	consonants := <-consonantCountCh

	fmt.Println("Count of wovel: ", vowels)
	fmt.Println("Count of consonant: ", consonants)
}
