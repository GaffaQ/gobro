package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, res := pembagian(100, 10)
	if res != nil {
		fmt.Println("Terjadi eror:", res)
	} else {
		fmt.Println("Hasilnya adalah:", hasil)
	}

	hasil2, res2 := pembagian(100, 0)
	if res2 != nil {
		fmt.Println("Terjadi eror:", res2)
	} else {
		fmt.Println("Hasilnya adalah:", hasil2)
	}
}
