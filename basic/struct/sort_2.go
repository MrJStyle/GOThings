package main

import (
	"fmt"
	"sort"
)

/*sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。
因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序。将代码1中的排序代码简化后如下所示：*/

func sortByStringSlice() {
	names := sort.StringSlice{
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

// 对结构体进行排序

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heroes []*Hero

func (s Heroes) Len() int {
	return len(s)
}

func (s Heroes) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}

	return s[i].Name < s[j].Name
}

func (s Heroes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	sortByStringSlice()

	heroes := Heroes{
		&Hero{"吕布", None},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Sort(heroes)

	for _, v := range heroes {
		fmt.Println(v)
	}
}
