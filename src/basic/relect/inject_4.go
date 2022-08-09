package main

import (
	"basic/relect/inject"
	"fmt"
)

type SpecialString2 interface{}
type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString2 `inject`
	uid    int            `inject`
	Age    int            `inject`
}

func main() {
	s := TestStruct{}
	inj := inject.New()
	inj.Map("张三")
	inj.MapTo("男", (*SpecialString2)(nil))
	inj2 := inject.New()
	// 可注入的前提是：字段必须是导出的（也即字段名以大写字母开头）
	inj2.Map(26)
	inj.SetParent(inj2)
	inj.Apply(&s)
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
}
