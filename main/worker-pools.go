package main

import (
	"fmt"
	"time"
)

/*
使用 goroutine 和通道来实现工作池。
*/

// 运行多个并发实例。这些工作人员将在 jobs 通道上接收工作，并在 results 上发送相应的结果。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	// 为了使用工作池，我们分别创建了两个通道用来发送数据和接收结果
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动3个工作协程
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 向工作池发送数据
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 关闭通道，表示已完成数据的发送
	close(jobs)

	// 等待协程执行完成，接收结果
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// 运行程序显示了不同工作人员正在执行的 5 个作业。
	// 尽管总共完成了大约 5 秒的工作，但该程序只花费了大约 2 秒，因为有 3 个工作人员同时运行。
}
