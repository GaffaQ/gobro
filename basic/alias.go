package main

import "fmt"

func main() {

	type NIM int

	var nim NIM = 25515040111015

	fmt.Println(nim)

	//===================

	type noktp string

	var nomer noktp = "12983112903"
	fmt.Println(nomer)

	num := noktp(nim)
	fmt.Println(num)

}
