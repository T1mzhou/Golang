package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 修改成员变量的值

// 接收者为普通变量，非指针，值语义，一份拷贝

// 带有接收者的函数叫方法

func (p *Person) SetInfoValue() {
	fmt.Println("SetInfoValue")

}

// 接受者为指针变量,引用传递
func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	// 结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	//p.SetInfoPointer() // func (p *Person) SetInfoPointer()
	(*p).SetInfoPointer() //把(*p)转换成p后再调用，等价于上面
	//内部做的转换， 先把指针p， 转成*p后再调用
	//(*p).SetInfoValue()
	//p.SetInfoValue()

}
