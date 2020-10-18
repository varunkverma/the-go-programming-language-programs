package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateByTwo(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
	fmt.Println()
	fmt.Println(b)
	rotateByTwo(b)
	fmt.Println(b)
}
