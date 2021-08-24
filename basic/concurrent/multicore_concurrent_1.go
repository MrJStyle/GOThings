package main

import (
	"fmt"
)

func AsyncFunc(index int, res chan int) {
	sum := 0
	for i := 0; i < 100000; i++ {
		sum += i
	}
	fmt.Printf("线程%d, sum为:%d\n", index, sum)
	res <- sum
}

func c1() {
	res := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go AsyncFunc(i, res)
	}

	for i := 0; i < 5; i++ {
		<-res
	}
}

func main() {
	c1()
}
