package main

import (
	"fmt"
	"strconv"
)

/*
从字符串中解析数字
内置包 strconv 提供数字解析。
*/

func main() {
	p := fmt.Println

	// 对于 ParseFloat ，这个 64 告诉要解析多少位精度。
	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	// 对于 ParseInt ， 0 表示从字符串推断基数。 64 要求结果适合 64 位。
	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	// ParseInt 将识别十六进制格式的数字。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	// ParseUint 也可用。
	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	// Atoi 是基本 10 int 解析的便捷函数。
	k, _ := strconv.Atoi("135")
	p(k)

	// 解析函数在输入错误时返回错误。
	_, e := strconv.Atoi("wat")
	p(e)
}
