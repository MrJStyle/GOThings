package main

import (
	"fmt"
	"reflect"
	"time"
)

func makeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		panic("expect a function")
	}

	wrapper := reflect.MakeFunc(tf, func(args []reflect.Value) (result []reflect.Value) {
		start := time.Now()
		result = vf.Call(args)
		end := time.Now()

		fmt.Printf("The Function takes %v\n", end.Sub(start))
		return result
	})

	return wrapper.Interface()
}

func timeMe() {
	time.Sleep(time.Second)

}

func timeMe2() bool {
	return true
}

func main() {
	timeFunc := makeTimeFunc(timeMe).(func())
	timeFunc()

	timeFunc2 := makeTimeFunc(timeMe2).(func() bool)
	fmt.Println(timeFunc2())
}
