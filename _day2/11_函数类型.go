package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

type FuncType func(int, int) int

func main() {
	var result int
	result = Add(1, 1)
	fmt.Println("result = ", result)

	var fTest FuncType
	fTest = Add
	result = fTest(10, 20)
	fmt.Print("result2 = ", result)

	fTest = Minus
	result = fTest(10, 5)
	fmt.Println("result3 = ", result)
}
