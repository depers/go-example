package main

import "fmt"

func main() {

	// 最基本的类型，具有单一条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始/条件/ for 之后循环。
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// 完成基本“执行 N 次”迭代的另一种方法是在整数上使用 range 。
	for i := range 3 {
		fmt.Println(i)
	}

	// 不带条件的 for 将重复循环，直到 break 跳出循环或 return 退出封闭函数。
	// 我理解就是while
	for {
		fmt.Println("loop")
		break
	}

	// continue的用法
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
