package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// read file
	// content,err := os.ReadFile(os.Args[1])
	content,err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error while reading the file", err)
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(log.Writer(),content)
}
