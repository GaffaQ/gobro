package main

import (
	"fmt"
	"time"
)

type Playerr struct {
	name       string
	registered time.Time
}

func (p Playerr) getRegistered() {
	fmt.Println(p.name)
	fmt.Println(p.registered.Year())
	fmt.Println(p.registered.Month())
	fmt.Println(p.registered.Day())
	fmt.Println(p.registered.Hour())
	fmt.Println(p.registered.Minute())
	fmt.Println(p.registered.Second())
}

func register(name string) *Playerr {
	return &Playerr{name, time.Now()}
}

func main() {

	gaffa := register("Gaffa")
	gaffa.getRegistered()
	time.Sleep(7 * time.Second)
	aini := register("Aini")
	aini.getRegistered()
}
