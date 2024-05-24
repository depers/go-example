package main

import "fmt"

/*
*

	关闭通道表示不会再在该通道上发送任何值。这对于向通道接收者传达完成情况很有用。
*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// 通道还没被关闭
			if more {
				fmt.Println("received job", j)
			} else {
				// 通道已经被发送者关闭
				// 通过 j, more := <-jobs 重复从 jobs 接收。在这种特殊的2个值形式的接收中，如果 jobs 已为 closed 并且所有值都在该频道已被接收。
				// 当我们完成所有工作时，我们用它来通知 done 。
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 通过 jobs 通道向工作线程发送 3 个作业，然后将其关闭。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 同步阻塞主线程，等待工作线程执行完毕
	<-done

	j, ok := <-jobs
	// 第一个参数：如果从关闭的通道上读取会返回0
	// 第二个参数：如果接收到的值是通过成功的发送操作传递到通道的，则可选的第二个返回值为 true ；
	//			 如果由于通道关闭且为空而生成零值，则可选的第二个返回值为 false 。
	fmt.Println("received more jobs:", j, ok)
}
