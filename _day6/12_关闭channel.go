package main

import (
	"fmt"
)

func main() {
	// 创建一个无缓存的channel
	ch := make(chan int)

	// 新建协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 往chan写内容
		}
		// 不需要写数据时，关闭channel
		close(ch)
		// ch <- 666 // 关闭channel后无法再发送数据
	}()

	for num := range ch {
		fmt.Println("num = ", num)
	}
}

func main01() {
	// 创建一个无缓存的channel
	ch := make(chan int)

	// 新建协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 往chan写内容
		}
		// 不需要写数据时，关闭channel
		close(ch)
		// ch <- 666 // 关闭channel后无法再发送数据
	}()

	for {
		// 如果ok为true，说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //  管道关闭
			break
		}
	}
}
