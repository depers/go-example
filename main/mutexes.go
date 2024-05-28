package main

import (
	"fmt"
	"sync"
)

/*
通过互斥锁来实现对跨域多个goroutine安全的访问数据
*/
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// 由于我们想要从多个 goroutine 并发更新它，因此我们添加 Mutex 来同步访问。
// 请注意，互斥体不能被复制，因此如果传递此 struct ，则应通过指针来完成。
func (c *Container) inc(name string) {
	// 在访问 counters 之前锁定互斥锁；
	c.mu.Lock()
	// 使用 defer 语句在函数末尾解锁它。
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// mu互斥量的零值可以按原样使用，不需要初始化
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	// 在循环中递增指定的计数器
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 同时运行多个 goroutine；请注意，它们都访问相同的 Container ，并且其中两个访问相同的计数器。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)

}
