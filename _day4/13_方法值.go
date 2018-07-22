package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)
	p.SetInfoPointer() // 传统方式

	// 保存方式入口地址
	pFunc := p.SetInfoPointer // 这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc()

	vFunc := p.SetInfoValue
	vFunc() // 等价于p.SetInfoValue()
}
