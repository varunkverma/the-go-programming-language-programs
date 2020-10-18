package main

import (
	"fmt"
)

func arrayDeclarationsAndInitialisations() {
	// array declaration:
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("at index %v, value: %v\n", i, v)
	}
	fmt.Println()
	// array initialisation
	var b [4]int = [4]int{1, 2, 3}
	for i, v := range b {
		fmt.Printf("at index %v, value: %v\n", i, v)
	}

	fmt.Println()

	c := [...]int{1, 2, 3}
	for i, v := range c {
		fmt.Printf("at index %v, value: %v\n", i, v)
	}

	//defines an array r with 100 elements, all zero except for the last, which has value −1.
	d := [...]int{99: -1}
	fmt.Println(d[0], d[50], d[99])
}

type currency int

// currencies of different types
const (
	USD currency = iota
	EUR
	GBP
	RMB
)

func symbolsAndCurrencies() {
	//it is also possible to specify a list of index and value pairs
	symbol := [...]string{USD: "$", EUR: " E ", GBP: " F ", RMB: " Y "}
	// remember USD = 0, EUR = 1 and so on
	fmt.Println(USD)
	fmt.Println(symbol[RMB])
}

func arrayCompare() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	//d := [3]int{1, 2}

	fmt.Printf("%v == %v : %v\n", a, b, a == b)
	fmt.Printf("%v == %v : %v\n", a, c, a == c)
	// fmt.Printf("%v == %v : %v", a, d, a == d) // invalid operation: a == d (mismatched types [2]int and [3]int)
}

func main() {
	arrayDeclarationsAndInitialisations()
	fmt.Println()
	symbolsAndCurrencies()
	fmt.Println()
	arrayCompare()
}
