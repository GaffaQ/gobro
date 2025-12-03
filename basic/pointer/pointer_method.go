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

func fungsiMenikah(man Man) {
	man.name = "Mr. " + man.name
}

func fungsiMenikahPointer(man *Man) {
	man.name = "Mr. " + man.name
}

func main() {

	gaffa := Man{"Gaffa"}
	gaffa.menikah()

	fmt.Println(gaffa)

	// karena dia menerima method dalam bentuk pointer
	gaffa.menikahPointer()
	fmt.Println(gaffa)

	yazid := Man{"Yazid"}
	fungsiMenikah(yazid)
	fmt.Println(yazid)
	fungsiMenikahPointer(&yazid)
	fmt.Println(yazid)

}
