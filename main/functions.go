package main

import "fmt"

// 这是一个函数，它接受两个 int 并将它们的总和作为 int 返回。
func plus(a int, b int) int {
	return a + b
}

// 当您有多个相同类型的连续参数时，您可以省略类似类型参数的类型名称，直到声明该类型的最后一个参数。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
