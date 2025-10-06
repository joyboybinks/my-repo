package main

import "fmt"

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w+h)
}
func main() {
	fmt.Println(RectPerimeter(3, 4))  // devrait afficher 14
	fmt.Println(RectPerimeter(5, 5))  // devrait afficher 20
	fmt.Println(RectPerimeter(-1, 2)) // devrait afficher -1
}