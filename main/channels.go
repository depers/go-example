package main

import "fmt"

/**
  Channels 是连接并发 goroutine 的管道。您可以将值从一个 Goroutine 发送到通道，并将这些值接收到另一个 Goroutine。
  默认情况下，发送和接收会阻塞，直到发送者和接收者都准备好为止。此属性允许我们在程序结束时等待 "ping" 消息，而无需使用任何其他同步。
*/

func main() {
	// 使用 make(chan val-type) 创建一个新通道。通道是根据它们传达的值来分类的
	message := make(chan string)

	// 使用 channel <- 语法将值发送到通道。在这里，我们从一个新的 goroutine 发送 "ping" 到我们上面创建的 messages 通道。
	go func() { message <- "ping" }()

	// <-channel 语法从通道接收一个值。在这里，我们将收到上面发送的 "ping" 消息并将其打印出来。
	msg := <-message
	// 当我们运行程序时， "ping" 消息通过我们的通道成功从一个 goroutine 传递到另一个 goroutine。
	fmt.Println(msg)
}
