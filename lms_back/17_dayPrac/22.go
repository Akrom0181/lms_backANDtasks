package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)


func IncrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	counter++
	fmt.Printf("Counter incremented to: %d\n", counter)
}

func main() {
	numGoroutines := 5

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go IncrementCounter(&wg)
	}
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
