package main

import "fmt"

func main() {

	//defer panicHandler()

	var x any = "apa"

	nilai, cek := x.(string)

	fmt.Println("nilai", nilai, "cek:", cek)

	if testPanic(false) {
		panic("PANIC WOI")
	}

	fmt.Println("===========================")

	safeAssertion(123)
	safeAssertion("Mantap")

}

func testPanic(cek bool) bool {
	return cek
}

func panicHandler() {
	msg := recover()
	fmt.Println("Recover", msg)
	fmt.Println("Panic berhasil di handle")
}

func safeAssertion(i any) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Hasil recover: ", msg)
		}
	}()

	val := i.(string)
	fmt.Println("Kalimat:", val)
}
