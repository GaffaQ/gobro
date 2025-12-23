package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name, text string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(text)
	return nil
}

func addToFile(name, text string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(text)
	return nil
}

func readFile(name string) (string, error) {

	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	msg := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		msg += string(line) + "\n"
	}

	return msg, nil

}

func main() {
	createNewFile("aini.txt", "Halo A ini gapaa\napa\n iyaa")
	msg, _ := readFile("aini.txt")
	fmt.Println(msg)
	addToFile("aini.txt", "\nGaffa sayang banget sama aii")
}
