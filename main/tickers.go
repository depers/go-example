package main

import (
	"fmt"
	"time"
)

/**
  股票行情记录器适用于您想要定期重复执行某项操作的情况
*/

func main() {

	// 每500ms执行一次
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			// 使用通道上的内置 select 来等待值每 500 毫秒到达一次。
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker将在1600毫秒后停止
	time.Sleep(1600 * time.Millisecond)

	// 在我们停止ticker之前，他会运行三次
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
