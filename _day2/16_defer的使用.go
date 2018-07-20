package main

import "fmt"

func main() {
	defer fmt.Println("bbbbbbbbbbb")

	fmt.Println("aaaaaaaaaaaa")
}
