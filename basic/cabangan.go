package main

import "fmt"

func main() {
	name := "a"

	if name == "a" {
		fmt.Println("a")
	} else if name == "gaffa" {
		fmt.Println("gaffa")
	} else {
		fmt.Println("beda")
	}

	// if short statement
	// panjang hanya berlaku di if statement ini saja, tidak berlaku diluar lagi.
	if panjang := len(name); panjang > 3 {
		fmt.Println("Panjang")
	} else {
		fmt.Println("Kurang")
	}

	// SWITCH EXPRESSION

	switch {
	case name == "a":
		fmt.Println("a")
	case name == "b":
		fmt.Println("b")
	case name == "c":
		fmt.Println("c")
	default:
		fmt.Println("boleh kenalan?")
	}
}
