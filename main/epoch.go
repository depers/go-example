package main

import (
	"fmt"
	"time"
)

/*
程序中的一个常见要求是获取自 Unix 纪元以来的秒数、毫秒数或纳秒数。
*/
func main() {
	now := time.Now()
	fmt.Println(now)

	// 将 time.Now 与 Unix 、 UnixMilli 或 UnixNano 一起使用可分别获取自 Unix 纪元以来经过的时间（以秒、毫秒或纳秒为单位）。
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 将自纪元以来的整数秒或纳秒转换为相应的 time
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
