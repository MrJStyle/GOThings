package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c  tac"

	reg1 := regexp.MustCompile("a.c")
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println(result1)
}
