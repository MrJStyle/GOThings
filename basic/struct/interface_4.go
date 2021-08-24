package main

import "fmt"

func main() {
	var a interface{} = 100

	var b interface{} = "aa"

	fmt.Println(a == b)
}
