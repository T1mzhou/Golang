package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("calc")
	result = fTest(a, b)

	return
}

func main() {
	a := Calc(1, 1, Mul)
	fmt.Println("a = ", a)
}
