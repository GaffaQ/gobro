package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	text := "Gaffa Fadhlanul Rozaq"
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Eror", err.Error())
	} else {
		//fmt.Println(decoded) dia bakal keluarnya slice bit, jadi di convert dlu biar dia jadi string
		fmt.Println(string(decoded))
	}

}
