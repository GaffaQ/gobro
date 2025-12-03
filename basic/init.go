package main

import (
	"gobro/basic/database"

	//buat nge init doang, tapi ga pake method di dalamnnya
	_ "gobro/basic/internal"
)

func main() {

	database.GetDatabase()

}
