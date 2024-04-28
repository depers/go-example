package main

import (
	"fmt"
)

/*
*

	数组
*/
func main() {
	// 创建一个数组 a ，它将恰好容纳 5 个 int 。元素的类型和长度都是数组类型的一部分。默认情况下，数组为零值，这对于 int 意味着 0
	var a [5]int
	fmt.Println("emp:", a)

	// 可以使用 array[index] = value 语法在索引处设置一个值，并使用 array[index] 获取一个值。
	a[4] = 100
	a[4] = 101
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置的 len 返回数组的长度
	fmt.Println("len:", len(a))

	// 可以在一行中声明和初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 可以使用 ... 让编译器为您计算元素数量
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println(b[0:2])

	// 使用 : 指定索引，则其间的元素将被清零
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// 可以组合类型来构建多维数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 可以立即创建和初始化多维数组
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
