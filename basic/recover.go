package main

import "fmt"

// defer panic recover = try catch
func end() {
	fmt.Println("end")
	msg := recover()
	fmt.Println(msg)
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
