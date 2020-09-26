package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	stop := time.Now()
	fmt.Println("For 1: ", stop.Sub(start))

	start = time.Now()
	for index, arg := range os.Args {
		println(index, arg)
	}
	stop = time.Now()
	fmt.Println("For 2: ", stop.Sub(start))
}
