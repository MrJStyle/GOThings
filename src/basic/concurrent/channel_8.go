package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg6 sync.WaitGroup

func init() {
	// 初始化随机种子
	rand.Seed(10)
}

func main() {
	tasks := make(chan string, taskLoad)

	wg6.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for t := 1; t <= taskLoad; t++ {
		tasks <- fmt.Sprintf("Task : %d", t)
	}

	// 当通道关闭后，goroutine依旧可以从通道接收数据，
	// 但是不能再像通道发送数据。
	// 能够从已经关闭的通道接收数据这一点非常重要， 因为这允许通道关闭后依旧能取出其中缓冲的全部值， 而不会有数据丢失
	close(tasks)

	wg6.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg6.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d shutting down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker %d start working: %s\n", worker, task)

		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示完成当前工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
