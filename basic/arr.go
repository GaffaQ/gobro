package main

import "fmt"

func main() {

	var players [3]string

	players[0] = "gaffa"
	players[1] = "baidu"
	players[2] = "baitu"

	fmt.Println(players)
	fmt.Println(players[0])
	fmt.Println(players[1])
	fmt.Println(players[2])

	//=======================

	var (
		names [2]string
		ages  [2]int
	)

	names[0] = "gaffa"
	names[1] = "baidu"

	ages[0] = 10
	ages[1] = 20
	fmt.Println(names)
	fmt.Println(ages)

	//=========================

	values := [3]int{
		1,
		2,
		3, //harus ada koma nya di akhir
	}
	fmt.Println(values)

	fal := [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(fal)
	fmt.Println(len(fal))
	fmt.Println(fal[len(fal)-1])

	//========================

}
