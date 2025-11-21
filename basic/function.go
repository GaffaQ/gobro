package main

import "fmt"

func main() {
	gapa()
	param("gaffa", 17)
	fmt.Println(sumall(1, 2))
	_, tipe := kategori(1, 2)
	fmt.Println(tipe)
	//fmt.Println(sum)
	fmt.Println(kategori(1, 2))

	fmt.Println(test())

	fmt.Println(sumALls(1, 2, 3, 4, 5))
}

// VOID FUNCTION
func gapa() {
	fmt.Println("Woi sigma")
}

func param(nama string, umur int) {
	fmt.Println("Halo", nama, "apakabar? umurlu", umur)
}

// RETURN VALUE FUNCTION
func sumall(a int, b int) int {
	return a + b
}

// RETURN MULTIVALUE FUNCTION
func kategori(a int, b int) (int, string) {

	sum := a + b
	tipe := ""
	if sum > 1 {
		tipe = "gacor"
	} else {
		tipe = "jelek"
	}

	return sum, tipe
}

func test() (first, mid, last string) {
	first = "apa"
	mid = "lu"
	last = "mau"

	return first, mid, last
}

// variadic function
func sumALls(numbers ...int) int {
	totals := 0
	for _, v := range numbers {
		totals += v
	}

	return totals
}
