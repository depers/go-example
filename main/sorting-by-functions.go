package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
按函数进行排序
*/
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 实现了字符串长度的比较函数
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		// 按年龄对 people 进行排序
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
