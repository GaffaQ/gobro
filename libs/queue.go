package main

import (
	"container/list"
	"fmt"
)

type queue struct {
	list *list.List
}

func NewQueuee() *queue {
	return &queue{list: list.New()}
}

func (q queue) Enqueue(v any) {
	q.list.PushBack(v)
}

func (q queue) Dequeue() {
	q.list.Remove(q.list.Front())
}

func (q queue) Len() int {
	return q.list.Len()
}

func (q queue) showData() {
	for i := q.list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func main() {

	antrian := NewQueuee()
	antrian.Enqueue("A")
	antrian.Enqueue("B")
	antrian.Enqueue("C")
	antrian.showData()
	fmt.Println(antrian.Len())
	antrian.Dequeue()
	antrian.showData()

}
