package main

import "fmt"

func findTwoSum(arr []int, total int) []int {
	cacheMap := map[int]int{}
	resArr := make([]int, 0)
	for idx, item := range arr {
		var another int = total - item
		if val, ok := cacheMap[another]; ok && val > 0 {
			resArr := append(resArr, item, another)
			fmt.Println(val, another)
			return resArr
		}
		cacheMap[item] = idx + 1
	}
	return resArr
}
func main() {
	inputArr := make([]int, 5)
	inputArr[0] = 2
	inputArr[1] = 7
	inputArr[2] = 11
	inputArr[3] = 5
	var target int = 9

	res := findTwoSum(inputArr, target)
	fmt.Println(res)
}
