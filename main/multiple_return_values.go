package main

import "fmt"

// 此函数签名中的 (int, int) 显示该函数返回 2个int
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 使用多重赋值使用2个不同的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 只需要返回值的子集，请使用空白标识符"_"
	_, c := vals()
	fmt.Println(c)
}
