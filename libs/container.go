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

	fmt.Println("============ MOVE ============")
}
