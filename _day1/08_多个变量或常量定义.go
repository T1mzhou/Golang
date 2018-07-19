package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2.0
	)
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	const (
		i = 10
		j = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
