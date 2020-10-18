package main

import (
	"fmt"
)

func main() {
	// printBytes("Hello World")
	// demoRawString()
	// demoUnicodeUTF8()
	// demoUnicodeRuneLiteral()

	str1 := "Hello, World"
	fmt.Println(hasPrefix(str1, "hello"))
	fmt.Println(hasSuffix(str1, "World"))
	fmt.Println(hasSubString(str1, ", W"))
}

func printBytes(s string) {
	fmt.Println(len(s))
	if len(s) > 0 {
		fmt.Println("s[0]:", s[0])
	}
}

func demoRawString() {
	const raw = `Ho Ho Ho
	Merry Christmas {1+1}\n`
	fmt.Println(raw)
}

func demoUnicodeUTF8() {
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")
}

func demoUnicodeRuneLiteral() {
	fmt.Println('\u4e16') //19990
}

func hasPrefix(s, prefix string) bool {
	return (len(s) >= len(prefix) && s[:len(prefix)] == prefix)
}

func hasSuffix(s, suffix string) bool {
	return (len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix)
}

func hasSubString(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
