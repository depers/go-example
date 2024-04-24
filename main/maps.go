package main

import (
	"fmt"
	"maps"
)

func main() {

	// 要创建空地图，请使用内置 make : make(map[key-type]val-type)
	m := make(map[string]int)

	// 使用典型的 name[key] = val 语法设置键/值对。
	m["k1"] = 7
	m["k2"] = 13
	// 使用例如打印地图 fmt.Println 将显示其所有键/值对。
	fmt.Println("m:", m)

	// 使用 name[key] 获取键的值。
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 如果键不存在，则返回值类型的零值。
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// 内置 len 返回键/值对的数量
	fmt.Println("len:", len(m))

	// 内置 delete 从映射中删除键/值对
	delete(m, "k2")
	fmt.Println("m:", m)

	// 要从映射中删除所有键/值对，请使用 clear 内置函数。
	clear(m)
	fmt.Println("m:", m)

	// 第一个值是value本身，第二值是该键是否存在于映射中
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 在同一行中声明并初始化一个新映射。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	// maps 包包含许多有用的地图实用函数
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
