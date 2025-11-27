package main

import "fmt"

func main() {

	//untuk menghentikan program
	cek(false)
	cek(true)

}

func setop() {
	fmt.Println("stop")
}

func middele() {
	fmt.Println("middle")
}

func cek(cek bool) {
	defer setop()
	if cek {
		panic("eror coy")
	}

	middele()
}
