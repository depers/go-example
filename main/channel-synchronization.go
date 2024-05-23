package main

import (
	"fmt"
	"time"
)

/**
  使用通道来同步各个 goroutine 的执行
*/

// 用阻塞接收来等待 goroutine 完成的例子
// 在 goroutine 中运行的函数。 done通道将用于通知另一个 goroutine 该函数的工作已完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完成
	done <- true
}

func main() {
	// 启动一个工作协程，为其提供通知通道
	done := make(chan bool, 1)
	go worker(done)

	// 阻塞直到我们收到通道上工作人员的通知。
	<-done
}
