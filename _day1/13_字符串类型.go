package main

import "fmt"

func main() {
	var str1 string
	str1 = "abc"
	fmt.Println("str1 = ", str1)

	str2 := "Tim"
	fmt.Printf("str2类型是 %T\n", str2)

	fmt.Println("len(str2) = ", len(str2))
}
