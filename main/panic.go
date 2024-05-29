package main

import "os"

/*
中止程序
常见用途是在函数返回我们不知道如何（或想要）处理的错误值时中止。
*/
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
