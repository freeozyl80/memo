package main

import "fmt"

type ListNode struct {
	value string
	next  *ListNode
}

func main() {
	var Map = map[int]bool{}
	Map[1] = true
	inputNumber := []int{1, 2, 3, 4}
	list1 := &ListNode{
		value: "123",
	}
	list2 := &ListNode{
		value: "1234",
	}

	fmt.Println(inputNumber)
	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(Map)
}
