package main

import "fmt"

// 定义一个Person结构体

type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建一个Person变量(对象)，这是个值类型，如果是引用类型注意要先make空间
	var person Person = Person{"jerry", 40}
	person.Age = 10
	fmt.Println(person)

	// 创建一个指针类型
	p := &Person{"a", 10}
	// 或
	// p := new(Person)
	m := p
	m.Age = 50
	fmt.Println(p)
}
