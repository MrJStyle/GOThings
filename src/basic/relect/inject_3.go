package main

import (
	"basic/relect/inject"
	"fmt"
	"reflect"
)

type SpecialString interface{}

func main() {
	inj := inject.New()
	inj.Map("C语言中文网")
	inj.MapTo("Golang", (*SpecialString)(nil))
	inj.Map(20)
	fmt.Println("字符串是否有效？", inj.Get(reflect.TypeOf("Go语言入门教程")).IsValid())
	fmt.Println("特殊字符串是否有效？", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
	fmt.Println("int 是否有效？", inj.Get(reflect.TypeOf(18)).IsValid())
	fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
	inj2 := inject.New()
	inj2.Map([]byte("test"))
	inj.SetParent(inj2)
	fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
}
