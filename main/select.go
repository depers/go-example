package main

import (
	"fmt"
	"time"
)

/**
  Go 的 select 允许您等待多个通道操作。
  这个操作类似于Java中提供的CountDownLatch
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 每个通道将在一段时间后收到一个值，以模拟例如阻塞在并发 goroutine 中执行的 RPC 操作。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 使用 select 同时等待这两个值，并在每个值到达时将其打印出来
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
