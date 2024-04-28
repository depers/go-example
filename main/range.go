package main

import "fmt"

/**
  关于切片和数组的范围属性
*/

func main() {

	// 使用 range 对切片中的数字求和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 数组和切片上的 range 提供每个条目的索引和值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// map上的 range 迭代键/值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 还可以仅迭代映射的键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 字符串上的 range 迭代 Unicode 代码点。第一个值是起始字节索引，第二个值是Unicode 代码点本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
