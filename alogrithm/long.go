package main

import "fmt"

// 在一个字符串重寻找没有重复字母的最长子串。
// awkeeee
func findMaxLen(input []rune) int {
	cache := make(map[rune]bool)
	max := 0
	length := len(input)
	i, j := 0, 0
	for i < length && j < length {
		if _, ok := cache[input[j]]; !ok {
			cache[input[j]] = true
			j++
			max = findMax(max, j-i)
		} else {
			delete(cache, input[i])
			i++
		}
	}

	return max
}
func findMax(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	inputStr := []rune("awkeeee")
	fmt.Println(findMaxLen(inputStr))

}
