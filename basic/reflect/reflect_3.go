package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 1024

	valueOfA := reflect.ValueOf(a)

	var getA = valueOfA.Interface().(int)

	var getB = int(valueOfA.Int())

	fmt.Println(getA, getB)
}
