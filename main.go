package main 

import "fmt"

func PrintIfNot(str string) string {
    if len(str) < 3 {
        return "G\n"
    }
    return "invalid Input\n"
}

func main() {
    fmt.Print(PrintIfNot("hi"))   // chaîne de 2 caractères → "G\n"
    fmt.Print(PrintIfNot("hello")) // chaîne de 5 caractères → "invalid Input\n"
}