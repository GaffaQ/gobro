package main

import "fmt"

type Man struct {
	name string
}

//tidak menerima pointer karena Man bukan *Man, artinya menduplikatnya
func (man Man) menikah() {
	man.name = "Mr. " + man.name
}

func (man *Man) menikahPointer() {
	man.name = "Mr. " + man.name
}

//function void with pointer

func fungsiMenikah(man Man) {
	man.name = "Mr. " + man.name
}

func fungsiMenikahPointer(man *Man) {
	man.name = "Mr. " + man.name
}

//return function with pointer
func createEntity(name string) *Man {
	return &Man{name}
}

func main() {

	//method with pointer
	gaffa := Man{"Gaffa"}
	gaffa.menikah()
	fmt.Println(gaffa)
	// karena dia menerima method dalam bentuk pointer
	gaffa.menikahPointer()
	fmt.Println(gaffa)

	//void function with pointer
	yazid := Man{"Yazid"}
	fungsiMenikah(yazid)
	fmt.Println(yazid)
	fungsiMenikahPointer(&yazid)
	fmt.Println(yazid)

	//return function with pointer
	fmt.Println(*createEntity("ripal"))

}
