package main

import "fmt"

func mufun01() (int, int, int) {
	return 1, 2, 3
}

// go官方推荐写法
func mufun02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}

func myfunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	a, b, c := mufun02()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
