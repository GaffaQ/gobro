package main

import (
	"fmt"
	"path"
)

func main() {

	//path
	fmt.Println(path.Dir("aini/gaffa.go"))             //dapetin directory ke filenya, return: aini
	fmt.Println(path.Base("aini/gaffa.go"))            //dapetin file targetnya, return: gaffa.go
	fmt.Println(path.Ext("aini/gaffa.go"))             //dapetin extensionnya, return .go
	fmt.Println(path.Join("gaffa", "aini", "mame.go")) //gabungin jadi gaffa/aini/mame.go

}
