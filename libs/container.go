package main

import (
	"container/list"
	"fmt"
)

func main() {

	//fmt.Println("hello world")
	data := list.New()
	data.PushBack("Depan")
	data.PushBack(2)

	data.PushBack("apa")

	//fmt.Println(data)

	fmt.Println("============ ITERATE ============")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		//fmt.Println(i)
	}

	fmt.Println("============ CMD ============")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Len())

	fmt.Println("============ INSERT ============")

	data2 := list.New()
	data2.PushBack("Satu")
	data2.PushBack("Tiga")

	data2.InsertAfter("Dua", data2.Front().Next())

	for i := data2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("============ MOVE ============")

	data3 := list.New()
	data3.PushBack("Satu")
	data3.PushBack("Tiga")
	data3.InsertBefore("Dua", data3.Back())
	data3.MoveToFront(data3.Back())

	for i := data3.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	data3.PushBack("Empat")
	data3.MoveAfter(data3.Back(), data3.Front())
	for i := data3.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("============ TYPE ASSERTION ============")

	data4 := list.New()

	data4.PushBack(1)
	data4.PushBack("Dua")
	data4.PushBack("Tiga")
	data4.PushBack("Empat")

	for i := data4.Front(); i != nil; i = i.Next() {

		val, ok := i.Value.(int)

		fmt.Println(val, ok)

	}

	fmt.Println("============ PRACTICE ============")
	fmt.Println("1. ")

	practice1 := list.New()
	practice1.PushBack(1)
	practice1.PushBack(2)
	practice1.PushBack(4)
	practice1.PushBack(5)

	for i := practice1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	practice1.InsertBefore(3, practice1.Front().Next().Next())
	for i := practice1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("2. ")
	practice2 := list.New()
	practice2.PushBack("A")
	practice2.PushBack("C")
	practice2.PushBack("B")
	practice2.PushBack("D")

	for i := practice2.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
