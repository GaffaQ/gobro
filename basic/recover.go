package main

import "fmt"

// defer panic recover = try catch
func end() {
	fmt.Println("end")
	msg := recover()
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Gada eror bamg")
	}
}

func siu(cek bool) {
	defer end()
	if cek {
		panic("EROR JIERRRR")
	}
}

func main() {

	siu(false)
	siu(true)

}
