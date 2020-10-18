package main

import "fmt"

func appendInt(oldSlice []int, element int) []int {
	var newSlice []int
	newSliceLen := len(oldSlice) + 1
	if newSliceLen <= cap(oldSlice) {
		// there's space in old_slice as its capacity is still >= length
		newSlice = oldSlice[:newSliceLen]
	} else {
		// There is insufficient memory in old_slice. Allocate new array.
		// Grow by doubling, fot amortized linear complexity
		newSliceCap := newSliceLen
		if newSliceCap < 2*len(oldSlice) {
			newSliceCap = 2 * len(oldSlice)
		}
		newSlice = make([]int, newSliceLen, newSliceCap)
		// copy data to new_slice from old_slice
		copy(newSlice, oldSlice)
	}
	newSlice[len(oldSlice)] = element
	return newSlice
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
