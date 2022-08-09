package main

import (
	"context"
	"fmt"
	"time"
)

func calcSomething(a int, ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Calculating")
		time.Sleep(time.Second)
	}
	ch <- a
}

func doSomethingWithContext(ctx context.Context) {
	ch := make(chan int)

	go calcSomething(10, ch)

	select {
	case x := <-ch:
		fmt.Printf("result %v", x)
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println("Timeout")
			fmt.Println(err.Error())
		}
	}
}

func main() {
	var ctx context.Context
	// 创建空context的两种方法
	ctx = context.Background() // 返回一个空的context，不能被cancel，kv为空

	//todoCtx := context.TODO()
	var cancel func()
	ctx, cancel = context.WithCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomethingWithContext(ctx)
}
