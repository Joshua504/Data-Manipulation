package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println("Error: file not found")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("can not close file", err)
			return
		}
	}(file)

	reader := csv.NewReader(file)
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		count++

		fmt.Printf("%-10s %-10s %-10s %-10s\n", record[0], record[1], record[2], record[3])
		if count == 1 {
			fmt.Println("---------- ---------- ---------- ----------")
		}
	}
}
