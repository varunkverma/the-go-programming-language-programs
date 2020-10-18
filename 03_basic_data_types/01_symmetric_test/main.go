package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5 // 0010 0010
	var y uint8 = 1<<1 | 1<<2 // 0000 0110

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)  // 0000 0010
	fmt.Printf("%08b\n", x|y)  // 0010 0110
	fmt.Printf("%08b\n", x^y)  // 0010 0100
	fmt.Printf("%08b\n", x&^y) /*
		x&(^y)
			0010 0010
		      & 1111 1001
			0010 0000

	*/
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i)
		}
	}
}
