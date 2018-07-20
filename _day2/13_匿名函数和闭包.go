package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	type FuncType func()

	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数，同时调研
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}()

	f3 := func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(1, 2)

	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(10, 20)

	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)

	fmt.Printf("x = %d, y = %d\n", x, y)
}
