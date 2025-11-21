package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("Selesai pertama\n")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Selesai kedua \n")

	iniSlice := make([]string, 2, 5)
	iniSlice[0] = "Gafa"
	iniSlice[1] = "Gafaa"

	for i, val := range iniSlice {
		fmt.Println(i, val)
	}

}
