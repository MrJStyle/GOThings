package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//读写锁有如下四个方法：
//写操作的锁定和解锁分别是func (*RWMutex) Lock和func (*RWMutex) Unlock；
//读操作的锁定和解锁分别是func (*RWMutex) Rlock和func (*RWMutex) RUnlock。
//
//读写锁的区别在于：
//当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
//当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；
//当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。
//
//所以说这里的读锁定（RLock）目的其实是告诉写锁定，有很多协程或者进程正在读取数据，写操作需要等它们读（读解锁）完才能进行写（写锁定）。
//
//我们可以将其总结为如下三条：
//同时只能有一个 goroutine 能够获得写锁定；
//同时可以有任意多个 gorouinte 获得读锁定；
//同时只能存在写锁定或读锁定（读和写互斥）。

var count int
var rw sync.RWMutex

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作\n", n)
	v := count
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d 读取结束， 值为%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束， 新值为%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
