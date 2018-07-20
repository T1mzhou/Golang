package main

import "fmt"

func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

func main() {
	defer fmt.Println("aaaaaaaaaaaaaaa")

	defer fmt.Println("bbbbbbbbbbbbbbb")

	//defer test(0)

	defer fmt.Println("cccccccccccccc")
}
