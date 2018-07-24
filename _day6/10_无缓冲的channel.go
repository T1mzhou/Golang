package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓存的channel
	ch := make(chan int, 0)

	// len(ch)缓冲区剩余数据个数， cap(ch)缓冲区的大小

	// 新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程: i = %d\n", i)
			ch <- i // 往chan写内容
		}
	}()

	//  延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch // 从管道中读取数据，没有内容之前是阻塞的
		fmt.Println("num = ", num)
	}
}
