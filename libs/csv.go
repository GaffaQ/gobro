package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//reader

	csvText := "Gaffa,Fadhlanul,Rozaq" +
		"Nur,Aini" +
		"Mima,Mame"

	reader := csv.NewReader(strings.NewReader(csvText))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	//writer

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Gaffa", "Fadhlanul", "Rozaq"})
	_ = writer.Write([]string{"Nur", "Aini"})
	_ = writer.Write([]string{"Mima", "Mame"})
	writer.Flush()
}
