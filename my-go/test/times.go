package main

import (
	"fmt"
)

// func findMaxTimes(inputStr string) {
// 	// 1. 遍历字符串，统计每个字符出现的次数
// 	charCount := make(map[rune]int)
// 	for _, char := range inputStr {
// 		charCount[char]++
// 	}

// 	// 2. 找出出现次数最多的字符
// 	maxChar := ' '
// 	maxCount := 0
// 	for char, count := range charCount {
// 		if count > maxCount {
// 			maxChar = char
// 			maxCount = count
// 		}
// 	}
// 	// 3. 打印出现次数最多的字符
// 	fmt.Println("出现次数最多的字符:", string(maxChar))

// 	// 4. 打印出现次数最多的字符的次数
// 	fmt.Println("出现次数:", maxCount)

// }

func findMaxTimes(inputStr string) {
	arr := []rune(inputStr)
	wordMap := make(map[string]int)
	for _, value := range arr {
		if wordMap[string(value)] == 0 {
			wordMap[string(value)] = 1
		} else {
			wordMap[string(value)]++
		}
	}
	fmt.Println(wordMap)
}

func main() {
	// inputStr1 := []rune("Hello世界") // unicode 数组
	// inputStr2 := []byte("Hello世界") // 字节数组
	// fmt.Println(len(inputStr1))
	// fmt.Println(len(inputStr2))
	// for _, value := range inputStr1 {
	// 	fmt.Println(string(value))
	// }
	// fmt.Println("Start")

	findMaxTimes("Hello世界")
}
