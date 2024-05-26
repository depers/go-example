package main

import "fmt"

/**
  可以使用for和range语法来迭代从通道接收的值。
*/

func main() {
	// 迭代 queue 通道中的 2 个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 这个例子还表明，可以关闭非空通道，但仍然可以接收剩余的值。
	close(queue)

	// range 会迭代从 queue 接收到的每个元素。因为我们 closed 了上面的通道，所以迭代在收到 2 个元素后终止。
	for elem := range queue {
		fmt.Println(elem)
	}
}
