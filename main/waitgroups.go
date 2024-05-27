package main

import (
	"fmt"
	"sync"
	"time"
)

/**
要等待多个 goroutine 完成，我们可以使用等待组。
*/

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {

	// WaitGroup用于等待这里启动的所有goroutines完成。
	// 注意：如果将 WaitGroup 显式传递给函数，则应通过指针来完成。
	var wg sync.WaitGroup

	// 启动多个 goroutine 并增加每个 goroutine 的 WaitGroup 计数器。
	for i := 1; i <= 5; i++ {
		// 启动多个 goroutine 并增加每个 goroutine 的 WaitGroup 计数器。
		wg.Add(1)

		// 将工作程序调用包装在一个闭包中，确保告诉 WaitGroup 该工作程序已完成。
		// 这样，工作线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			// 这个就是闭包中的外部状态
			// 在worker函数执行完之后执行
			defer wg.Done()
			// 这个是函数逻辑
			worker(i)
		}()
	}

	// 阻塞直到WaitGroup计数器回到0
	// 请注意，此方法没有直接的方法来传播工作函数的错误。对于更高级的用例，请考虑使用 errgroup 包。
	// 每次调用的工作程序启动和完成的顺序可能不同。
	wg.Wait()
}
