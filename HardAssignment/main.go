package main

import (
	"fmt"
	"io"

	// "log"
	"os"
)

func main() {

	fileName := os.Args[1]

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// data, err := os.ReadFile(file.Name())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf(string(data))

	io.Copy(os.Stdout, file)

	file.Close()

}
