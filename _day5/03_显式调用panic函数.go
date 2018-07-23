package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaa")
}
func testb() {
	// fmt.Println("bbbbbbbbb")
	// 显式调用panic函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("ccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
