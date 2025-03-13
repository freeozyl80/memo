package main

import "fmt"

//	给定 {a{bc{d}}}
//
// 输出 d bc a

func parser(input string) string {
	arr := []rune(input)

	stack := []string{}

	for _, item := range arr {
		// 如果当前字符是 '{'，则在栈中添加一个空字符串
		if string(item) == "{" {
			stack = append(stack, "")
			// 如果当前字符是 '}'，则将栈顶的字符串合并到前一个字符串中
		} else if string(item) == "}" {
			if len(stack) > 1 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1] += top
			}
			// 如果当前字符是其他字符，则将其添加到栈顶的字符串中
		} else {
			stack[len(stack)-1] += string(item)
		}
	}

	result := ""
	for i := len(stack) - 1; i >= 0; i-- {
		result += stack[i] + " "
	}

	return result[:len(result)-1]
}

func main() {
	fmt.Println("开始")
	// input := make(map[string]string)
	// input["test"] = "123"
	// fmt.Println(input)
	res := parser("{a{bc{d}}}")
	fmt.Println(res)

	arr := [][]int{
		{1, 2, 3},
	}
	fmt.Println(arr)
}
