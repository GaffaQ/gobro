package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	data := ring.New(5)

	for i := 1; i <= data.Len(); i++ {
		data.Value = "Halo " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(v any) {
		fmt.Println(v)
	})

}
