package main

import "fmt"

/*
可以使用任意数量的尾随参数来调用可变参数函数
*/

// 这个函数接受任意数量的 int 作为参数
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	// 在函数内， nums 的类型相当于 []int 。我们可以调用 len(nums) ，使用 range 迭代它，等等
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	// 可变参数函数可以通过单独的参数以通常的方式调用
	sum(1, 2)
	sum(1, 2, 3)

	// 如果切片中已经有多个参数，请使用 func(slice...) 将它们应用到可变参数函数
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
