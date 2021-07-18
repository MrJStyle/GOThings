package main

import (
	"basic/relect/inject"
	"fmt"
)

type S3 interface {
}

type S4 interface {
}

type Staff struct {
	Name    string `inject`
	Company S3     `inject`
	Level   S4     `inject`
	Age     int    `inject`
}

func main() {
	// 创建注入实例
	s := Staff{}
	// 控制实例的创建
	inj := inject.New()
	// 初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S3)(nil))
	inj.MapTo("T4", (*S4)(nil))
	inj.Map(23)
	// 实现对struct注入
	inj.Apply(&s)
	// 打印结果
	fmt.Printf("s = %v\n", s)
}
