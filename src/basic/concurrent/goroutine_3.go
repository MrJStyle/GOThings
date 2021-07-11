package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&count, 1) //安全的对counter加1
		runtime.Gosched()
	}
}
