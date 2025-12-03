package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() {
	fmt.Println(connection)
}
