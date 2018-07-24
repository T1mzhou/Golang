package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个有缓存的channel
	ch := make(chan int, 3)

	// len(ch)缓冲区剩余数据个数， cap(ch)缓冲区的大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	// 新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 往chan写内容
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	//  延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch // 从管道中读取数据，没有内容之前是阻塞的
		fmt.Println("num = ", num)
	}
}
