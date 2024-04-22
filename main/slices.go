package main

import (
	"fmt"
	"slices"
)

func main() {

	// 未初始化的切片等于 nil，长度为 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// 创建长度非零的空切片，请使用内置 make 。这里我们制作了一个长度为 3 的 string 切片（最初为零值）。默认情况下，新切片的容量等于其长度
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// 可以像数组一样设置和获取
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 按预期返回切片的长度
	fmt.Println("len:", len(s))

	// 一种是内置的 append ，它返回一个包含一个或多个新值的切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片的复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片支持语法为 slice[low:high] 的“切片”运算符。例如，这会获取元素 s[2] 、 s[3] 和 s[4] 的切片
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 切片到下标5（但不包括） s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// (包括） s[2] 切分出来
	l = s[2:]
	fmt.Println("sl3:", l)

	// 在一行中声明并初始化切片的变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// slices 包包含许多有用的切片实用函数
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 切片可以组成多维数据结构。与多维数组不同，内部切片的长度可以变化
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
