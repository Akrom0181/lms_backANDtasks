package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	numGoroutines := 5

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			if err := Do(); err != nil {
				fmt.Println("Error:", err)
			}
		}()
	}

	wg.Wait()
}

func Do() error {
	if randomError() {
		return errors.New("error occured: ")
	}
	return nil
}

func randomError() bool {
	return true 
}
