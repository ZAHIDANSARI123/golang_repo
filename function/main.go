package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learing function in Golang")
	simpleFunction()

	ans := add(3, 6)
	fmt.Println("add of two number is :", ans)

	data := multiply(7, 8)
	fmt.Println("multiplication of two number is :", data)
}
