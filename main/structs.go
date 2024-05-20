package main

import "fmt"

// person 结构类型具有 name 和 age 字段
type person struct {
	name string
	age  int
}

// 构造函数，使用newPerson函数构造一个新的 person 结构
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// 创建一个新结构
	fmt.Println(person{"Bob", 20})

	// 在初始化结构时命名字段
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将为零值
	fmt.Println(person{name: "Fred"})

	// & 前缀产生指向该结构的指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 新结构体的创建封装在构造函数中是惯用的做法
	fmt.Println(newPerson("Jon"))

	// 使用点访问结构体字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 可以将点与结构指针一起使用，指针会自动取消引用
	sp := &s
	fmt.Println(sp.age)

	// 结构是可变的
	sp.age = 51
	fmt.Println(sp.age)

	// 如果结构类型仅用于单个值，则不必为其指定名称。该值可以具有匿名结构类型。此技术通常用于表驱动测试。
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
