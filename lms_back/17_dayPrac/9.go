package main

import (
	"fmt"
	"sort"
)

func sorting(numbers []int, ch chan []int) {
	sort.Ints(numbers)
	ch <- numbers
}

func main()  {
	nums := []int{9,3,5,3,6,72,7,9,2,4,5}
	ch := make(chan []int)
	fmt.Println("before sorting:", nums)

	go sorting(nums, ch)

	sorted := <-ch
	fmt.Println("after sorted: ", sorted)
}
