package main

import "fmt"

type Player struct {
	name, address string
	age           int
}

func main() {
	var player Player
	player.name = "Alex"
	player.address = "US"
	player.age = 18
	fmt.Println(player)

	var player1 Player
	player1.name = "Agus"
	player1.address = "USA"
	player1.age = 19
	fmt.Println(player1)

	//cara lain
	gaffa := Player{
		name:    "Gaffa",
		address: "ID",
		age:     18,
	}

	fmt.Println(gaffa)

	//atau
	aini := Player{"Aini", "ID", 19}
	fmt.Println(aini)
}
