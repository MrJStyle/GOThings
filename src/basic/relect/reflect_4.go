package main

import (
	"fmt"
	"reflect"
)

// 普通函数

func Add(a, b int) int {
	return a + b

}

func main() {
	// 将函数包装为反射对象
	methodV := reflect.ValueOf(Add)

	// 构造函数参数，并传入两个值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	// 反射调用函数
	resList := methodV.Call(paramList)

	// 获取第一个返回值，取整型
	fmt.Println(resList[0].Int())
}
