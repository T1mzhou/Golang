package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
	name   string // 与Person同名了
}

func main() {
	// 声明（定义一个变量）
	var s Student
	// 默认规则（纠结原则），如果能在本作用域找到此成员，
	// 就操作此成员如果没有找到，找到继承的字段
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	// 显式调用
	s.Person.name = "yoyo" // Person的name
	fmt.Printf("s = %+v\n", s)
}
