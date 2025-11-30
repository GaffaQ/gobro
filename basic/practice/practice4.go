package main

import "fmt"

type Pointer struct {
	name string
}

func (p *Pointer) setName(name string) {
	p.name = name
}

func setName(name string, p *Pointer) {
	p.name = name
}

func main() {

	person := Pointer{"Gaffa"}
	fmt.Println(person.name)

	person.setName("Aini")
	fmt.Println(person.name)

	setName("Rozaq", &person)
	fmt.Println(person.name)

	sendMessage("wooi")
	sendMessage(12)
	sendMessage(true)
}

func sendMessage(teks any) {
	fmt.Println(teks)
}
