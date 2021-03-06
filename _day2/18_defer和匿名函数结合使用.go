package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Printf("a = %d, b= %d\n", a, b)
	}(10, 20)

	a = 111
	b = 222
	fmt.Printf("外部: a = %d, b = %d\n", a, b)
}

func main01() {
	a := 10
	b := 20
	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部: a = %d, b = %d\n", a, b)
}
