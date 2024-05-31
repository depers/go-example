package main

import (
	"fmt"
	"os"
)

/*
字符串格式化
*/
type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// 打印结构体
	fmt.Printf("struct1: %v\n", p) // struct1: {1 2}

	fmt.Printf("struct2: %+v\n", p) // struct2: {x:1 y:2}

	fmt.Printf("struct3: %#v\n", p) // struct3: main.point{x:1, y:2}

	// 打印值的类型
	fmt.Printf("type: %T\n", p) // type: main.point

	// 打印布尔值
	fmt.Printf("bool: %t\n", p) // bool: {%!t(int=1) %!t(int=2)}

	fmt.Printf("int: %d\n", 123)

	// 二进制表示
	fmt.Printf("bin: %b\n", 14)

	// 给定整数对应的ascii字符
	fmt.Printf("char: %c\n", 33)

	// 16进制编码
	fmt.Printf("hex: %x\n", 456)

	// 浮点数
	fmt.Printf("fliat1: %f\n", 78.9)

	// 使用科学计数法格式的浮点数
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	// 以16为基数显示字符串
	fmt.Printf("str3: %x\n", "hex this")

	// 打印指针
	fmt.Printf("pointer: %p\n", &p)

	// 指定宽度且右对齐
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 限制小数精度
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 左对齐
	fmt.Printf("witdh3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 字符串右对齐
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 字符串左对齐
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf 格式化并返回字符串，而不在任何地方打印它。
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// 使用 Fprintf 格式化+打印到 io.Writers 而不是 os.Stdout 。
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
