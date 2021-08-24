package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj interface{}) {
	// pair: <value, type>
	t := reflect.TypeOf(obj)
	n := t.Name()

	// A Kind represents the specific kind of type that a Type represents.
	// The zero Kind is not a valid kind.
	k := t.Kind()

	v := reflect.ValueOf(obj)

	fmt.Printf("Type: %s\nType.Name: %s\nKind: %s\nValue: %v\n\n", t, n, k, v)
}

type handler func(int, int) int

func main() {
	var intVar int64 = 10
	stringVar := "Hello"
	type book struct {
		name  string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}

	var sub handler = func(a, b int) int {
		return a - b
	}

	sli := make([]int, 5)

	printMeta(intVar)
	printMeta(stringVar)
	printMeta(book{
		name:  "harry potter",
		pages: 500,
	})

	printMeta(sum)
	printMeta(sub)
	printMeta(sli)
}
