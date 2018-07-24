package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccc")

	// return // 终止此函数
	runtime.Goexit() // 终止所在的协程
	fmt.Println("ddddddddddd")
}

func main() {
	// 创建新建的协程
	go func() {
		fmt.Println("aaaaaaaaaaaa")

		// 调用别的函数
		test()

		fmt.Println("bbbbbbbbbbbbb")
	}() // 别忘了()

	// 特地写一个死循环，目的不让主协程结束
	for {

	}
}
