package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 只有类型，没有名字，匿名字段，继承了Person的成员
	int    // 基础类型的匿名字段
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "hehehe"}
	fmt.Printf("s = %+v\n", s)

	s.Person = Person{"go", 'm', 22}
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person, s.int, s.mystr)
}
