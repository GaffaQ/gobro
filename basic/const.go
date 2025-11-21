package main

import "fmt"

func main() {

	//const = gabisa diubah lagi setelah di declare (FINAL)

	const nim int = 25515040111015
	const nama = "Gaffa"

	const (
		firstName = "Gaffa"
		lastName  = "Fadhlanul"
	)

	//eror
	//nim = 1233123
	//nama = "sigma"

	fmt.Println(nama)
	fmt.Println(nim)

	fmt.Println(lastName)
	fmt.Println(firstName)

}
