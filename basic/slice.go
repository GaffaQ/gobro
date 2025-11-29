package main

import "fmt"

func main() {

	names := [...]string{
		"gaffa",
		"fadhlanul",
		"rozaq",
		"joko",
	}

	fmt.Println(names)

	fmt.Println(names[1:])  // [fadhlanul rozaq joko]
	fmt.Println(names[:3])  // [gaffa fadhlanul rozaq]
	fmt.Println(names[0:2]) // [gaffa fadhlanul]
	fmt.Println(names[1:3]) // [fadhlanul rozaq]

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days) //otomatis keubah juga, karena dari slicenya berubah.

	// APPEND menambah data, maksimal mencapai kapasitas, jika len nya 2, tpai capasitasnya 5, berarti bisa menambahkan 3 nilai lagi

	daySlice2 := append(daySlice1, "hari baru", "apa")
	fmt.Println(daySlice2)

	play := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	playSlice := play[1:3]
	fmt.Println(playSlice)
	playSlice[0] = 5
	fmt.Println(playSlice)
	fmt.Println(play)
	playSlice2 := append(playSlice, 7, 6, 5)
	fmt.Println(playSlice2)
	fmt.Println(playSlice)
	fmt.Println(play)

	//MAKE membuat slice baru dengan array baru

	newSlice := make([]string, 2, 5)
	fmt.Println(newSlice)
	newSlice[0] = "Gaffa"
	newSlice[1] = "Fadhlanul"
	fmt.Println(newSlice)
	newSlice2 := append(newSlice, "Sabtu baru")
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
	fmt.Println("cap", cap(newSlice), "len", len(newSlice), "cap", cap(newSlice2), len(newSlice2))
	newSlice2[0] = "sigma"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//perulangan
	slicegw := []string{"gaffa", "fadhlanul", "rozaq", "jumat", "sabtu"}
	for i := 0; i < len(slicegw); i++ {
		fmt.Println(slicegw[i])
	}

	for id, val := range slicegw {
		fmt.Println(id, val)
	}
}
