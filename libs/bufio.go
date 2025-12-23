package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//read

	text := "Gaffa Sayang Aini" + "Aini Sayang Gaffa"
	reader := bufio.NewReader(strings.NewReader(text))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	//writer

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Gaffa Sayang Aini")
	_, _ = writer.WriteString("Aini Sayang Gaffa")
	writer.Flush()
}
