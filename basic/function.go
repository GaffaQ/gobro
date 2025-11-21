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

	// anonymous function

	banned := func(name string) bool {
		if name == "anjing" {
			return true
		}
		return false
	}

	regist("anjing", banned)
	regist("Gaffa", func(name string) bool {
		if name == "anjing" {
			return true
		}
		return false
	})

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

type fungsi func(string) string

func getAnimal(name string) string {
	return name + "!"
}

func helloAnimal(name string, animal fungsi) string {
	return "You animal name is " + animal(name)
}

//func helloAnimal(name string, animal func(string) string) string {
//	return "You animal name is " + animal(name)
//}

// anonymous function

type Banneds func(string) bool

func regist(name string, banned Banneds) {
	if banned(name) {
		fmt.Println("Banned!", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
