package main

import "fmt"

func main() {
	var a int = 10
	// 每个变量有2层含义：变量对的内存,变量的地址
	fmt.Printf("a = %d\n", a) // 变量的内存
	fmt.Printf("&a = %v\n", &a)

	// 保存某个变量的地址,需要指针类型, *int保存int类变量的地址
	var p *int
	p = &a // 指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 666 // *p操作的不是p的内存，是p所指向的内存（就是a）
	fmt.Printf("*p = %v, a = %v\n", *p, a)

}
