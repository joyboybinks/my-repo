package main

import "fmt"

func PrintIf(str string) string {
	if len(str) <= 3 || str == "" {
		return "G\n"
	}
	return "Invalid Invalid\n"
}

func main() {
	fmt.Print(PrintIf("hi"))     // → "G\n"
	fmt.Print(PrintIf(""))       // → "G\n"
	fmt.Print(PrintIf("test"))   // → "Invalid Invalid\n"
}