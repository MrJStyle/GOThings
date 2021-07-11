package main

import (
	"fmt"
)

func PrintType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	PrintType(111)
	PrintType("hello")
	PrintType(true)
	PrintType(111.5)
}
