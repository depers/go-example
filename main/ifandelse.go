package main

import (
	"fmt"
)

func main() {

	// 基本示例
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 不带 else 的 if 语句
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// && 和 || 等逻辑运算符在条件中通常很有用
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// 语句可以位于条件语句之前；该语句中声明的任何变量都可以在当前分支和所有后续分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// go中是没有三元运算符的，就是 int a = 1 > 0 ? 2 : 3
}
