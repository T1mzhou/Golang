package main

import "fmt"

func myfunc01() int {
	return 666
}

func myfunc02() (result int) {
	return 666
}

func myfunc03() (result int) {
	result = 666
	return
}

func main() {
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc01()
	fmt.Println("b = ", b)

	c := myfunc03()
	fmt.Println("c = ", c)
}
