package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
原子计数器
*/
func main() {

	// 使用原子整数类型来表示我们的（始终为正）计数器。
	var ops atomic.Uint64

	// WaitGroup 将帮助我们等待所有 goroutine 完成其工作。
	var wg sync.WaitGroup

	// 我们将启动 50 个 goroutine，每个 goroutine 都会将计数器精确增加 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// 为了以原子方式递增计数器，我们使用 Add 。
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	// 等待所有 goroutine 完成。
	wg.Wait()

	// 使用 Load 可以安全地原子读取一个值，即使其他 goroutine 正在（原子地）更新它。
	fmt.Println("ops:", ops.Load())
}
