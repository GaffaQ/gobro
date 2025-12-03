package main

import "fmt"

type Person struct {
	name string
	age  int
}

func ubahNama(nama *Person) {
	nama.name = "apalu"
}

func ubahNilai(nilai *int) {
	*nilai = 10
}

func main() {

	a := 10
	b := &a //pointer
	fmt.Println(a, *b)
	*b = 100
	fmt.Println(a, *b)

	var person Person = Person{"Aini", 20}
	var person2 *Person = &person

	person2.age = 19
	fmt.Println(person, *person2)

	fmt.Println("=============================")

	// Arterisk operator *
	person3 := Person{"Gaffa", 20}
	person4 := &person3
	fmt.Println(person, person4)
	fmt.Println("=============================")

	orang := Person{"Oke gas", 19}
	fmt.Println(orang.name)
	ubahNama(&orang)
	fmt.Println(orang.name)

	fmt.Println("=============================")

	nilaibaru := 100
	fmt.Println(nilaibaru)
	ubahNilai(&nilaibaru)
	fmt.Println(nilaibaru)

	fmt.Println("=============================")

}
