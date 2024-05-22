package main

import "fmt"

/**
  通道缓冲
  默认情况下，通道是无缓冲的
*/

func main() {

	//  make 一个字符串通道，最多缓冲 2 个值。
	messages := make(chan string, 2)

	// 由于该通道是缓冲的，因此我们可以将这些值发送到通道中，而无需相应的并发接收。
	messages <- "buffered"
	messages <- "channel"

	// 接收这两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
