package main

import "fmt"

type Rectangel struct {
	name string
	s    int
}

type Square struct {
	name string
	p, l int
}

type Shape interface {
	Area() int
	getName() string
}

func (square Square) getName() string {
	return square.name
}

func (rect Rectangel) getName() string {
	return rect.name
}

func (square Square) Area() int {
	return square.p * square.l
}

func (rect Rectangel) Area() int {
	return rect.s * rect.s
}

func getLuas(intface Shape) {
	fmt.Println("Luas", intface.getName(), "adalah :", intface.Area())
}

func main() {
	Persegi := Rectangel{"Persegi", 10}
	persegiPanjang := Square{"Persegi Panjang", 10, 5}
	getLuas(Persegi)
	getLuas(persegiPanjang)
}
