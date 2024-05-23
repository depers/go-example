package main

import "fmt"

/**
  在使用通道作为函数参数时，可以指定通道是用于接收还是发送，从而提高程序的类型安全性
*/

// 此 ping 函数仅接受用于发送值的通道。尝试在此通道上接收将是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接受一个用于接收的通道 ( pings )，另一个用于发送的通道 ( pongs )。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// 将msg发送到pings通道
	ping(pings, "passed message")
	// 从pings取出msg，然后将msg再发送到pongs
	pong(pings, pongs)
	// 接收pongs的值
	fmt.Println(<-pongs)
}
