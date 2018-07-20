package main

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func myfunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func test(args ...int) {
	myfunc(args...)

	myfunc2(args[:2]...)
	myfunc2(args[2:]...)
}

func main() {
	test(1, 2, 3, 4)
}
