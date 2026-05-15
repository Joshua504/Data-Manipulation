package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type database struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type config struct {
	AppName  string   `json:"app_name"`
	Version  string   `json:"version"`
	Port     int      `json:"port"`
	Debug    bool     `json:"debug"`
	Database database `json:"database"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage go run . [filename]")
		return
	}

	open, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: file does not exist")
		return
	}

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(open)

	var config config

	dec := json.NewDecoder(open)

	err = dec.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AppName: %s\nVersion: %s\nPort: %d\nDebug: %t\nDB-Host: %s\n", config.AppName, config.Version,
		config.Port, config.Debug, config.Database.Host)
}
