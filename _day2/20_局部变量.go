package main

import "fmt"

func test() {
	a := 10
	fmt.Println("a = ", a)
}

func main() {
	{
		i := 10
		fmt.Println("i = ", i)
	}

	if flag := 3; flag == 3 {
		fmt.Println("flag = ", flag)
	}
	flag = 4
}
