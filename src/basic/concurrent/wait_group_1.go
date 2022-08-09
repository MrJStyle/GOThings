package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// 声明一个等待组
	var wg1 sync.WaitGroup

	var urls = []string{
		"http://www.baidu.com/",
		"https://www.qq.com/",
		"https://www.hao123.com",
	}

	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时，将等待组增加1
		wg1.Add(1)

		// 开启一个并发
		go func(url string) {
			defer wg1.Done()

			_, err := http.Get(url)

			fmt.Println(url, err)
		}(url)
	}

	wg1.Wait()

	fmt.Println("over")
}
