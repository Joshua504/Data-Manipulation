package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: not enough argument")
		return
	}
	dir := os.Args[1]

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error: can not read file", err)
		return
	}

	fmt.Printf("%-20s %-10s %-20s\n", "Name", "Size", "Last Modified")
	fmt.Println("--------------------  ----------  --------------------")

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("Error: can not provide info on this file", err)
			return
		}
		fmt.Printf("%-10s %-10d %-10s\n", info.Name(), info.Size(), info.ModTime().Format("2006-01-02 15:04"))
	}
}
