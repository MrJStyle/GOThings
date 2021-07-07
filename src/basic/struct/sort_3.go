package main

import (
	"fmt"
	"sort"
)

type HeroKind1 int

const (
	None1 = iota
	Tank1
	Assassin1
	Mage1
)

type Hero1 struct {
	Name string
	Kind HeroKind1
}

func main() {
	heros := []*Hero1{
		{"吕布", Tank1},
		{"李白", Assassin1},
		{"妲己", Mage1},
		{"貂蝉", Assassin1},
		{"关羽", None1},
		{"诸葛亮", Mage1},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
