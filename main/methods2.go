package main

import "fmt"

// 值接收器和指针接收器的区别
// 如果结构体方法的接收器是值类型，则在方法内使用的结构体是原有结构体拷贝的副本
// 如果结构体方法的接收器是指针类型，则在方法内使用的结构体是原有的结构体

type Person struct {
	name string
	age  int
}

func (p *Person) printInfo() {
	p.name = "Alice"

}

func (p Person) printValue() {
	p.name = "Bob"
}

func main() {
	// 调用者p1是值类型
	p1 := Person{name: "Tom", age: 25}

	// 调用指针接收器方法
	p1.printInfo()
	fmt.Println("name:", p1.name)
	fmt.Println("age:", p1.age)

	// 调用值接收器方法
	p1.printValue()
	fmt.Println("name:", p1.name)
	fmt.Println("age:", p1.age)

	// 调用者p2是指针类型
	p2 := &Person{name: "小张", age: 18}
	// 调用值接收器方法
	p2.printValue()
	fmt.Println("name:", p2.name)
	fmt.Println("age:", p2.age)

	// 调用指针接收器方法
	p2.printInfo()
	fmt.Println("name:", p2.name)
	fmt.Println("age:", p2.age)

}
