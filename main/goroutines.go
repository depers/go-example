package main

import (
	"fmt"
	"time"
)

/**
  Goroutine 是一个轻量级的执行线程
*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 同步执行
	f("direct")

	// 要在 goroutine 中调用此函数，请使用 go f(s) 。
	// 这个新的 goroutine 将与调用的 goroutine 同时执行
	go f("goroutine")

	// 启动一个 goroutine 来进行匿名函数调用
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 这里对goroutine异步运行进行等待
	time.Sleep(time.Second)
	fmt.Println("done")
}
