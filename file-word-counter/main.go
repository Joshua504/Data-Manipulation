package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	rFile, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error: file does not exist", err)
		return
	}

	defer func(rFile *os.File) {
		err := rFile.Close()
		if err != nil {

		}
	}(rFile)

	scanner := bufio.NewScanner(rFile)

	count, countChar, countStr := 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		countChar += utf8.RuneCountInString(line)

		countStr += len(strings.Fields(line))

		//* counting lines
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: reading file", err)
		return
	}

	fmt.Printf("Line: %d\n", count)
	fmt.Printf("Character: %d\n", countChar)
	fmt.Printf("Word: %d\n", countStr)
}

//func countcharAndStr(a, b, c string) int {
//
//	for range a{
//		a++
//	}
//	for range b{
//		b++
//	}
//
//}
