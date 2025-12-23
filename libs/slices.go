package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"Gaffa", "Aini"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(names, "Gaffa"))
	fmt.Println(slices.Index(names, "Aini"))

}
