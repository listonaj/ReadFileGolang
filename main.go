package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Replace "yourfile.txt" with the path to your text file
	filePath := "SNOWY-EVENING.txt"

	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Display the content
	fmt.Println("File Content:")
	fmt.Println(string(content))
}
