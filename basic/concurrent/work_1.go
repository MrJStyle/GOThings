package main

import (
	"basic/concurrent/work"
	"log"
	"sync"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现Worker接口
func (n *namePrinter) Task() {
	log.Println(n.name)
	//time.Sleep(time.Duration(100) * time.Millisecond)
}

func main() {
	// 使用两个goroutine来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup

	times := 1000000
	wg.Add(times * len(names))

	for i := 0; i < times; i++ {
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行。当Run返回时我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，并等待所有现有的工作完成
	p.Shutdown()
}
