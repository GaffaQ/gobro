package main

import "fmt"

func main() {

	start()
	defer stop()
	middle()

}

func stop() {
	fmt.Println("stop")
}

func start() {
	fmt.Println("start")
}

func middle() {
	fmt.Println("middle")
}
