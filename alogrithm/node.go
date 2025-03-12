package main

import "fmt"

// Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
// Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
type Node struct {
	value int
	next  *Node
}
type ListNode struct {
	list *[]Node
}

func main() {
	inputNode1 := &Node{
		value: 9,
		next:  nil,
	}
	inputNode2 := &Node{
		value: 1,
		next:  inputNode1,
	}

	input1 := &ListNode{}
	var list []Node
	list = append(list, *inputNode1, *inputNode2)
	input1.list = &list
	fmt.Println(input1)
}
