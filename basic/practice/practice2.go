package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func variadic(angka ...int) int {
	sum := 0
	for _, v := range angka {
		sum += v
	}
	return sum
}

func functionAsValue(nama string) {
	fmt.Println("Halo", nama)
}

func functionAsValue2(nama string) string {
	return "Haloo " + nama
}

type leng func(strings) int

func funcAsParam(nama string, fung leng) {
	fmt.Println("Halo", nama, fung(nama))
}

func main() {

	fmt.Println(variadic(1, 2, 3))

	func1 := functionAsValue
	func2 := functionAsValue2
	func1("aini")
	fmt.Println(func2("gapa"))

	func3 := func(a string) int {
		return len(a)
	}
	func4 := funcAsParam
	func4("Gaffa", func3)

}
