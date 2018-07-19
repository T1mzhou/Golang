package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	fmt.Println("t = ", t)
}
