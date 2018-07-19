package main

import "fmt"

func main() {
	str := "abc"

	for i := 0; i < len(str); i++ {
		fmt.Println("str[%d]=%c\n", i, str[i])
	}

	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}
	for i := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
