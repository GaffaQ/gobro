package main

import "fmt"

func main() {

	var angka int
	// angka := 10
	const PI = 4.14
	fmt.Println(angka)

	angka1 := 1928.12
	fmt.Println(int(angka1), angka1)

	//teks := "gaffa"

	sum := ""
	slice1 := []string{"gaffa", "suka", "ai"}
	for i := range slice1 {
		sum += slice1[i]
	}
	fmt.Println(sum)

	//map
	map1 := map[string]int{
		"gaffa": 1,
		"suka":  2,
		"ai":    3,
	}
	sum2 := 0
	for _, v := range map1 {
		sum2 += v
	}
	fmt.Println(sum2)

}
