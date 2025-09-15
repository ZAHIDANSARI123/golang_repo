package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the first statement")
	data := add(5, 10)
	fmt.Println("sum is :", data)
	defer fmt.Println("This is the second statement")
	fmt.Println("This is the third statement")
}

func add(a int, b int) int {
	return a + b
}
