package main

import (
	"fmt"
	"sort"
)

type Player struct {
	name  string
	level int
}

type Players []Player

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i, j int) bool {
	return p[i].level > p[j].level
}

func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	players := []Player{
		{"Gaffa", 100},
		{"Aini", 101},
		{"Mame", 90},
	}
	sort.Sort(Players(players))
	
	for _, v := range players {
		fmt.Println(v.level)
	}
}
