package main

import (
	"fmt"
	"strings"
)

func main() {
	// "hellogo"中是否包含"hello", 包含返回true， 不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	// Joins组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println("buf = ", buf)

	// Index，查找子串位置
	fmt.Println(strings.Index("abcdhello", "hello"))
	fmt.Println(strings.Index("abcdhello", "go")) // 不包含子串返回-1

	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf) // gogogo

	// Split 以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	// Trim去掉俩头的字符
	buf = strings.Trim("           are u ok?             ", " ")
	fmt.Printf("buf = #%s#\n", buf)

	// 去掉空格，把元素放入切片
	s3 := strings.Fields("   are u ok ?        ")
	// fmt.Println("s3 = ", s3)
	for i, data := range s3 {
		fmt.Println(i, ",", data)
	}
}
