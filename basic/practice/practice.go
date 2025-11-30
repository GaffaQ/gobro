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
	fmt.Println("ok", sum2)

	//buatkan array
	arr1 := [3]int{1, 2, 3}
	for _, k := range arr1 {
		fmt.Println(k)
	}

	//buat slicer

	slice2 := [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	fmt.Println(slice2)

	slice3 := slice2[:3]
	slice4 := append(slice3, 10, 11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println(slice4)
	fmt.Println(slice3)

	slice5 := slice4[0:3]
	slice5[0] = 100
	slice4[1] = 31
	fmt.Println(slice5)
	fmt.Println(slice4)
	fmt.Println(slice3)
	fmt.Println(slice2)

	slice6 := make([]int, 5, 5)
	fmt.Println(slice6)
	slice6 = append(slice6[:1], slice6[2:]...)
	slice6[0] = 10
	slice6[1] = 11
	slice6[2] = 12
	fmt.Println(slice6)

	slice7 := make([]int, 2, 5)
	slice7[0] = 1
	slice7[1] = 2
	fmt.Println(slice7)
	slice8 := copy(slice7, slice6)
	fmt.Println("delapan", slice8)

	// map
	map2 := map[string]int{
		"gaffa": 19,
		"ai":    20,
	}

	for _, k := range map2 {
		fmt.Println(k)
	}
	delete(map2, "gaffa")

	//array 2d
	arr2d := [2][2]int{
		{1, 1},
		{2, 2},
	}
	fmt.Println(arr2d)

	arr5 := [3]int{1, 2, 3}
	slicearr := arr5[:]
	slicearr[0] = 100
	fmt.Println(arr5)
	fmt.Println(slicearr)

	user1 := User{"gaffa", 12}
	slicestruct := make([]User, 10)
	fmt.Println(slicestruct)
	slicestruct[0] = user1
	fmt.Println(slicestruct)
	for _, v := range slicestruct {
		fmt.Println(v.name, v.age)
	}

}

type User struct {
	name string
	age  int
}
