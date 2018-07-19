package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	a = 8
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

	if a := 8; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a <10")
	} else {
		fmt.Println("这是不可能的")
	}
}
