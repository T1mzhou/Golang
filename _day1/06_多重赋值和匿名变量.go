package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a, b := 10, 20

	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b = %d\n", a, b)

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)
	i = 10
	j = 20

	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var c, d, e int
	c, d, e = test()
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	_, d, e = test()
	fmt.Printf("d = %d, e = %d\n", d, e)
}
