package main

import (
	"fmt"
	"sync"
)

func sum(start, end int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	resultChan <- sum
}

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan int, 3)

	wg.Add(3)
	go sum(1, 33, &wg, resultChan)
	go sum(34, 66, &wg, resultChan)
	go sum(67, 100, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for result := range resultChan {
		totalSum += result
	}

	fmt.Println("The sum from 1 to 100 is:", totalSum)
}
