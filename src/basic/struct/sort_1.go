package main

import (
	"fmt"
	"sort"
)

// 对一系列字符串进行排序时，使用字符串切片（[]string）承载多个字符串。使用 type 关键字，
//将字符串切片（[]string）定义为自定义类型 MyStringList。为了让 sort 包能识别 MyStringList，
//能够对 MyStringList 进行排序，就必须让 MyStringList 实现 sort.Interface 接口。

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Println(v)
	}
}
