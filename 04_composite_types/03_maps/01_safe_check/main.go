package main

import (
	"fmt"
)

func checkIfKeyPresentExample() {
	name := make(map[string]int)
	_, ok := name["alice"]
	if !ok {
		fmt.Println("Name not present")
	} else {
		fmt.Println("name present")
	}
}

func main() {
	checkIfKeyPresentExample()
}
