package main

import (
	"fmt"
	"sort"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(value int) {
	// newNode := &Node{value: value}
	//  m := make(map[string]int)
	newNode := new(Node)
	newNode.value = value
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Display()

	arr := []int{1, 2, 12, 3}
	arr = append(arr, 1)
	sort.Ints(arr)
	fmt.Println(arr)
}
