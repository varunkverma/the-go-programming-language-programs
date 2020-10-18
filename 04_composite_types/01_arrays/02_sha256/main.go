package main

import (
	"crypto/sha256"
	"fmt"
)

func compareStringsUsingHash(a, b string) {
	h1 := sha256.Sum256([]byte(a))
	h2 := sha256.Sum256([]byte(b))

	fmt.Printf("%x\n%x\n%t\n%T\n", h1, h2, h1 == h2, h1)
}

func main() {
	compareStringsUsingHash("varun", "Varun")
	fmt.Println()
	compareStringsUsingHash("varun", "varun")

}
