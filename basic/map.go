package main

import "fmt"

func main() {

	iniMap := make(map[int]string)
	// iniMap := map[int]string {    }

	iniMap[1] = "gaffa"
	iniMap[2] = "jumat"

	fmt.Println(iniMap)

	delete(iniMap, 1)
	fmt.Println(iniMap)

	mapgue := map[int]string{
		0: "satu",
		1: "dua",
	}

	for i := 0; i < len(mapgue); i++ {
		fmt.Println(mapgue[i])
	}

	mapkedua := make(map[string]int)
	mapkedua["satu"] = 1
	mapkedua["dua"] = 2
	fmt.Println(mapkedua)

	for k, v := range mapkedua {
		fmt.Println(k, v)
	}

}
