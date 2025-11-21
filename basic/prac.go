package main

import "fmt"

func main() {

	//1
	slicegw := []int{3, 5, 7, 2}
	sum := 0
	for i := 0; i < len(slicegw); i++ {
		sum += slicegw[i]
	}
	fmt.Println(sum)

	//2
	maks := 0
	slicekedua := []int{4, 6, 1, 9, 17}
	for i := 0; i < len(slicekedua); i++ {
		if slicekedua[i] > maks {
			maks = slicekedua[i]
		}
	}
	fmt.Println(maks)

	//3
	sliceketiga := []int{1, 2, 3, 4}
	reversed := make([]int, len(sliceketiga))

	for i := range reversed {
		reversed[i] = sliceketiga[len(sliceketiga)-i-1]
	}
	fmt.Println(reversed)

	//4
	slicekeempat := []int{1, 2, 3, 4, 5, 6}
	var genap []int

	for _, v := range slicekeempat {
		if v%2 == 0 {
			genap = append(genap, v)
		}
	}
	fmt.Println(genap)

	//5
	slicekelima := []int{1, 2, 1, 3, 1}
	target := 1
	cnt := 0

	for _, v := range slicekelima {
		if v == target {
			cnt++
		}
	}
	fmt.Println(cnt)

	//6
	slicekeenam := []int{1, 2, 2, 3, 3, 3}
	var dupe []int
	for _, v := range slicekeenam {
		flag := true
		for _, w := range dupe {
			if v == w {
				flag = false
				break
			}
		}

		if flag {
			dupe = append(dupe, v)
		}
	}
	fmt.Println(dupe)

	//7
	sliceketujuh := []int{10, 20, 30, 40}

	f := sliceketujuh[0]
	for i := 0; i < len(sliceketujuh)-1; i++ {
		sliceketujuh[i] = sliceketujuh[i+1]
	}
	sliceketujuh[len(sliceketujuh)-1] = f
	fmt.Println(sliceketujuh)

	//8

	// prac function
	sliceee := []int{1, 2, 3, 4, 8}
	fmt.Println(getMaks(sliceee...))

	fungsi1 := func(s int) int {
		return s * s
	}

	fmt.Println(calc(5, fungsi1))

	multip := multiple(5)
	fmt.Println(multip(3))

	fmt.Println(join(".", "a", "k", "c", "inta", " ", "aiii"))

}

func getMaks(numbers ...int) int {

	maks := numbers[0]
	for _, v := range numbers {
		if v > maks {
			maks = v
		}
	}

	return maks

}

type calcu func(int) int

func calc(s int, ans calcu) int {
	return ans(s)
}

func multiple(s int) func(int) int {
	return func(n int) int {
		return s * n
	}
}

func join(sep string, words ...string) string {

	wordss := "" + sep
	for _, v := range words {
		wordss += v + sep
	}

	return wordss
}
