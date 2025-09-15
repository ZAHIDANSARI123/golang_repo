package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	/**
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	content := "Hello, World!"
	bytes, err := file.WriteString(content + "\n")
	fmt.Println("Bytes written:", bytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created successfully")

	**/

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// create a buffer to hold the file content

	buffer := make([]byte, 100)

	// read the file into the buffer
	for {
		n, err := file.Read(buffer)
		if err != io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// proccess the read bytes
		fmt.Println(string(buffer[:n]))
	}

	contact, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(contact))
}
