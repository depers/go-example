package main

import (
	"fmt"
	"time"
)

/*
速率限制，我猜底层是通过令牌桶算法做的实现
速率限制是控制资源利用率和维持服务质量的重要机制。 Go 优雅地支持 goroutine、通道和代码的速率限制。
*/
func main() {
	// 先在通道中放5个请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 速率控制器，每隔200ms接收一个值
	limiter := time.Tick(200 * time.Millisecond)

	// 限制为每 200 毫秒 1 个请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 允许短时间内的请求突发，同时保持总体的速率限制
	// 通过缓冲限制器通道来实现，此 burstyLimiter 通道最多允许突发 3 个事件。
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道以表示允许的突发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每隔 200 毫秒，我们就会尝试向 burstyLimiter 添加一个新值，最多可达 3。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 现在再模拟 5 个传入请求。其中前 3 个将受益于 burstyLimiter 的突发功能。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burst request", req, time.Now())
	}
}
