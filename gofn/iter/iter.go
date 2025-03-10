package main

import (
	"fmt"
	"strings"
)

// 这里 type parameters 的概念 范型
func Convert[S any, D any](src []S, mapFn func(s S) D) []D {
	r := make([]D, 0, len(src))
	for _, i := range src {
		r = append(r, mapFn(i))
	}
	return r
}

func ToUpByString() {
	sl := []string{"hello", "world", "golang"}
	s0 := Convert(sl, func(v string) string { return strings.ToUpper(v) })
	for _, v := range s0 {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Hello, playground")
	ToUpByString()
}
