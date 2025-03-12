package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type hashMap struct {
	size   int
	bucket []*Node
}

func main() {
	var Map = map[int]int{}
	Map[1] = 1
	inputNumber := []int{1, 2, 3, 4}
	list1 := &Node{
		value: "123",
	}
	list2 := &Node{
		value: "1234",
	}
	inputNumber = append(inputNumber, 1, 2)
	fmt.Println(inputNumber)
	fmt.Println(list1)
	fmt.Println(list2)
	item := Map[10]
	fmt.Println(item)
	delete(Map, 1)
	fmt.Println(Map)
}
