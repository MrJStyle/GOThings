package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter2 int
	wg5      sync.WaitGroup
	mutex    sync.Mutex
)

func incCounter() {
	defer wg5.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter2

			runtime.Gosched()

			value++

			counter2 = value
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	wg5.Add(2)

	go incCounter()
	go incCounter()

	wg5.Wait()
	fmt.Printf("Final Counter value is %d", counter2)
}
