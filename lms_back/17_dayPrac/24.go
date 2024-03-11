package main

import (
	"fmt"
	"sync"
)

func main() {
	const numberOfWorkers = 5
	const numberOfTasks = 10

	tasks := make(chan int, numberOfTasks)
	results := make(chan int, numberOfTasks)

	var wg sync.WaitGroup

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(tasks, results, &wg)
	}

	for i := 0; i < numberOfTasks; i++ {
		tasks <- i
	}

	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}
}

func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		result := task * 2
		results <- result
	}
}
