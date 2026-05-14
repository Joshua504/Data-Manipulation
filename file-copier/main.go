package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: not enough argument")
		return
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: file not found", err)
		return
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("can not close file", err)
			return
		}
	}(src)

	dst, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error: file not be created", err)
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			fmt.Println("can not close file", err)
			return
		}
	}(dst)

	byteCopied, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("can not copy content from %d to %d", src, dst)
		log.Fatal(err)
	}
	log.Printf("Successfully copied %d bytes", byteCopied)
}
