package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := os.Args[1]

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(data))

	file.Close()

}
