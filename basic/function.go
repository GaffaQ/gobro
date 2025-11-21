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

	fmt.Println(sumAllslice([]int{1, 2, 3, 4, 5}))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumALls(nums...))

	hello := getGreeting
	fmt.Println(hello("Gaffa"))

	fmt.Println(helloAnimal("Anjing", getAnimal))
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

func sumAllslice(numbers []int) int {

	totals := 0
	for _, v := range numbers {
		totals += v
	}
	return totals
}

// function as value
func getGreeting(name string) string {
	return "Hello " + name
}

// parameter as function
func getAnimal(name string) string {
	return name + "!"
}

func helloAnimal(name string, animal func(string) string) string {
	return "You animal name is " + animal(name)
}
