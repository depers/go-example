package main

import (
	"fmt"
	"time"
)

/*
*

	定时器,适用于您想要在将来执行一次某项操作的情况
*/
func main() {

	fmt.Println(time.Now())
	// 计时器在等待2秒后，会提供一个通道C
	timer1 := time.NewTimer(2 * time.Second)

	// 主线程阻塞在通道C上2秒
	time1 := <-timer1.C
	fmt.Println("Timer 1 fired", time1)

	// 计时器在等待1秒后执行
	timer2 := time.NewTimer(time.Second)
	// 在协程中接收timer2通道的值
	go func() {
		// 在协程中阻塞
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// 关闭计时器2
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 主线程给 timer2 足够的时间来触发
	time.Sleep(2 * time.Second)
}
