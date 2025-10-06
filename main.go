package main 

func PrintIfNot(str string) string {
    if len(str) < 3 {
        return "G\n"
    }
    return "invalid Input\n"
}