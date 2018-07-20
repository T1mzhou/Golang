package main

import "fmt"

var a byte

func main() {
	var a int

	fmt.Printf("%T\n", a)
	{
		var a float32
		fmt.Printf("2:%T\n", a)
	}
	test()
}

func test() {
	fmt.Printf("3: %T\n", a)
}
