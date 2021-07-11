package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)

			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println("show input: ", input)
}
