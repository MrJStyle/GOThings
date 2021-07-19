package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) setAge(age int) {
	s.Age = age
}

func (s Student) GetAge() int {
	return s.Age
}

func test(a, b int) bool {
	return false
}

func main() {
	s1 := Student{
		Name: "Tom",
		Age:  10,
	}
	st := reflect.TypeOf(s1.setAge)
	fmt.Println(st.In(0))
}
