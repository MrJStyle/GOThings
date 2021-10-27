package main

import (
	"fmt"
	"reflect"
)

func main() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("is float canSet %v, CanAddr %v\n", v.CanSet(), v.CanAddr())
	//v.SetFloat(222222.211)    // 会报错

	vp := reflect.ValueOf(&fVar)
	fmt.Printf("is float canSet %v, CanAddr %v\n", vp.Elem().CanSet(), vp.Elem().CanAddr())
	vp.Elem().SetFloat(2.333333)

	println(fVar)
}
