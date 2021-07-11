package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg2      sync.WaitGroup
)

func doWork(name string) {
	defer wg2.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s work\n", name)
			break
		}
	}
}

func main() {
	wg2.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(time.Second * 2)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg2.Wait()
}
