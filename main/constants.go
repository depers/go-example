package main

import (
	"fmt"
	"math"
)

/*
 声明常量
*/

// const语句可以出现在var语句可以出现的任何位置。
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000
	// 常量表达式以任意精度执行算术
	const d = 3e20 / n
	fmt.Println(d)

	// 数字常量在被赋予类型之前没有类型，例如通过显式转换。
	fmt.Println(int64(d))

	// 在变量赋值或是函数调用的时候，会自动对常量赋予类型
	fmt.Println(math.Sin(n))

}
