package main

import "fmt"

func main() {
	fmt.Println("4 > 3结果:", 4 > 3)
	fmt.Println("4 > 3结果:", 4 != 3)
	fmt.Println("4 > 3结果:", !(4 > 3))
	fmt.Println("4 > 3结果:", !(4 != 3))

	fmt.Println("true && true 结果:", true && true)
	fmt.Println("true &&| flalse 结果:", true || false)

	fmt.Println("true || false 结果: ", true || false)
	fmt.Println("false || false 结果: ", false || false)

	a := 11
	fmt.Println("0 <= a && a<= 10结果为: ", 0 <= a && a <= 10)
}
