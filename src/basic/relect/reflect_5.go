package main

import (
	"fmt"
	"reflect"
)

func InvertSlice(args []reflect.Value) (result []reflect.Value) {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {

	invert := reflect.ValueOf(p).Elem()

	//Use of MakeFunc() method
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

// Main function
func main() {

	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{1, 2, 3, 4, 2, 3, 5}))

}
