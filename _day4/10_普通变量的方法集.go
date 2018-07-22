package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) SetInfoValue() {
	fmt.Println("SetInfoValue")

}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	p := Person{"mike", 'm', 18}
	p.SetInfoPointer() // func (p *Person) SetInfoPointer()
	// 内部，先把p, 转为为&p再调用， (&p).SetInfoPointer()
	p.SetInfoValue()
}
