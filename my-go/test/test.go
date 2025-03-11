package main

import (
	"fmt"
	"sync"
)

func appendToslice(res *[]int, arr []int, i int, wg *sync.WaitGroup) {

	defer wg.Done()
	*res = append(arr, i)
	for i, _ := range arr {
		arr[i] += 100
	}
}

func main() {
	var res []int
	arr := []int{1, 2, 3}
	var wg sync.WaitGroup
	wg.Add(2)
	go appendToslice(&res, arr, 1, &wg)
	go appendToslice(&res, arr, 3, &wg)
	wg.Wait()
	fmt.Println(res)
}
