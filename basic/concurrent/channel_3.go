//无缓冲通道

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg3.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won.\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed.\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit.\n", name)
		ball++

		court <- ball
	}
}

func main() {
	// 创建一个无缓冲通道
	court := make(chan int)

	// 计数加2， 表示要等待两个goroutine
	wg3.Add(2)

	go player("Ada", court)
	go player("Steve", court)

	time.Sleep(time.Second)
	court <- 1

	wg3.Wait()
}
