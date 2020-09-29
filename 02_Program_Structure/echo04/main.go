package main

import (
	"flag"
	"fmt"
	"strings"
)

var seperator = flag.String("s", " ", "seperator")
var shouldOmitNextline = flag.Bool("n", false, "omit trailing newline")

func main() {
	flag.Parse() // update the flag variables from their default values.
	fmt.Print(strings.Join(flag.Args(), *seperator))
	if !*shouldOmitNextline {
		fmt.Println()
	}
}
