package calc

import "fmt"

func init() {
	fmt.Println("This is calc init")
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
