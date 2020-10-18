package main

import (
	"fmt"
)

func removeElementFromIndex(numbers []int, index int) []int {
	copy(numbers[index:], numbers[index+1:])
	return numbers[:len(numbers)-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)
	nums = removeElementFromIndex(nums, 4)
	fmt.Println(nums)
}
