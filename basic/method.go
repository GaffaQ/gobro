package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (Person Person) say() {
	fmt.Println("the moon is beautiful, isn't it?", Person.name)
}

func main() {
	aini := Person{"Aini", 20}
	aini.say()
}
