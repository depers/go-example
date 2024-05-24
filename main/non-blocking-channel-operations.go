package main

import "fmt"

/**
  通道上的基本发送和接收是阻塞的。
  但是，我们可以使用 select 和 default 子句来实现非阻塞发送、接收，甚至非阻塞多路 select 。
*/

func main() {
	messages := make(chan string)
	singals := make(chan bool)

	// 非阻塞接收。如果 messages 上有可用值，则 select 将采用具有该值的 <-messages case 。
	// 如果不是，它将立即处理 default 情况。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送的工作原理类似。这里 msg 无法发送到 messages 通道，因为该通道没有缓冲区，也没有接收器。
	// 因此选择 default 情况。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 子句上方使用多个 case 来实现多路非阻塞选择。
	// 在这里，我们尝试在 messages 和 signals 上进行非阻塞接收。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-singals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
