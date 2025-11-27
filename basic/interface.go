package main

import "fmt"

type Name interface {
	getName() string
}

type Hewan struct {
	name string
	foot int
}
type Human struct {
	name string
	age  int
}

func sayHello(intface Name) {
	fmt.Println("Hello", intface.getName())
}

func (hewan Hewan) getName() string {
	return hewan.name
}

func (human Human) getName() string {
	return human.name
}

func main() {

	aini := Human{"Aini", 20}
	michi := Hewan{"Michi", 4}

	sayHello(aini)
	sayHello(michi)

}
